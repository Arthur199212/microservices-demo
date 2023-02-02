package data

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetSupportedCurrencies(t *testing.T) {
	supportedCurrencyCodes := []string{"EUR", "USD", "GBP"}
	cd := &currencyData{
		rates:                  make(map[string]float32),
		supportedCurrencyCodes: supportedCurrencyCodes,
	}
	currencies := cd.GetSupportedCurrencies()
	assert.Len(t, currencies, len(supportedCurrencyCodes))
	assert.Equal(t, supportedCurrencyCodes, currencies)
}

func TestMonitorRates(t *testing.T) {
	callsNum := 0
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cubes := Cubes{
			CubeData: []Cube{
				{
					Currency: "USD",
					Rate:     "1.1",
				},
				{
					Currency: "GBP",
					Rate:     "1.2",
				},
			},
		}
		data, err := xml.Marshal(cubes)
		assert.NoError(t, err)

		callsNum++

		w.Header().Add("Content-Type", "application/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}))

	ecbURL = svr.URL
	cd, err := NewCurrencyData()
	assert.NoError(t, err)

	cd.MonitorRates(time.Second)

	expectedCallsNum := 4
	time.Sleep(time.Duration(expectedCallsNum) * time.Second)

	assert.Equal(t, expectedCallsNum, callsNum)
}

func TestVerifySupportedCurrency(t *testing.T) {
	cd := &currencyData{
		rates: map[string]float32{
			"EUR": 1,
			"USD": 1.1,
			"GBP": 1.2,
		},
	}

	err := cd.VerifySupportedCurrency("EUR")
	assert.NoError(t, err)
	err = cd.VerifySupportedCurrency("XXX")
	assert.Error(t, err)
}

func TestConvert(t *testing.T) {
	cd := &currencyData{
		rates: map[string]float32{
			"EUR": 1,
			"USD": 1.1,
			"GBP": 1.2,
		},
	}

	res, err := cd.Convert("USD", "EUR", 100)
	assert.NoError(t, err)
	assert.Equal(t, float32(90.91), res)

	res, err = cd.Convert("USD", "GBP", 50)
	assert.NoError(t, err)
	assert.Equal(t, float32(54.55), res)

	_, err = cd.Convert("XXX", "EUR", 100)
	assert.Error(t, err)

	_, err = cd.Convert("USD", "XXX", 100)
	assert.Error(t, err)
}
