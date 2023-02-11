package data

import (
	"encoding/xml"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type CurrencyData interface {
	Convert(fromCurrency, toCurrency string, amount float32) (float32, error)
	GetSupportedCurrencies() []string
	// Checks the rates in the ECB API every interval.
	MonitorRates(interval time.Duration)
	VerifySupportedCurrency(currencyCode string) error
}

type currencyData struct {
	mu                     sync.RWMutex
	rates                  map[string]float32
	supportedCurrencyCodes []string
}

func NewCurrencyData() (CurrencyData, error) {
	cd := &currencyData{
		mu:                     sync.RWMutex{},
		rates:                  make(map[string]float32),
		supportedCurrencyCodes: []string{},
	}
	err := cd.getRates()
	return cd, err
}

func (cd *currencyData) MonitorRates(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C
			log.Info().Msg("running monitor rates job")
			cd.getRates()
		}
	}()
}

func (cd *currencyData) VerifySupportedCurrency(currencyCode string) error {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	_, ok := cd.rates[currencyCode]
	if !ok {
		return fmt.Errorf("currencyCode=%s is not supported", currencyCode)
	}
	return nil
}

func (cd *currencyData) Convert(
	fromCurrency string,
	toCurrency string,
	amount float32,
) (float32, error) {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	fromRate, ok := cd.rates[fromCurrency]
	if !ok {
		return 0, fmt.Errorf("rate not found for currencyCode=%s", fromCurrency)
	}
	toRate, ok := cd.rates[toCurrency]
	if !ok {
		return 0, fmt.Errorf("rate not found for currencyCode=%s", toCurrency)
	}
	raw := amount / fromRate * toRate
	truncated := math.Round(float64(raw)*100) / 100 // formats to '%.2f'
	return float32(truncated), nil
}

func (cd *currencyData) GetSupportedCurrencies() []string {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	return cd.supportedCurrencyCodes
}

var ecbURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

// Gets the rates from ECB (European Central Bank)
func (cd *currencyData) getRates() error {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	res, err := http.DefaultClient.Get(ecbURL)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected http code=200 but got code=%d", res.StatusCode)
	}
	defer res.Body.Close()

	cubes := &Cubes{}
	if err := xml.NewDecoder(res.Body).Decode(&cubes); err != nil {
		return fmt.Errorf("could not decode fetched data: %v", err)
	}

	for _, cube := range cubes.CubeData {
		rate, err := strconv.ParseFloat(cube.Rate, 32)
		if err != nil {
			return err
		}
		cd.rates[cube.Currency] = float32(rate)

		cd.supportedCurrencyCodes = append(cd.supportedCurrencyCodes, cube.Currency)
	}

	// EUR is the base rate
	cd.rates["EUR"] = 1
	cd.supportedCurrencyCodes = append(cd.supportedCurrencyCodes, "EUR")

	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
