/*
IMPORTANT: 	THIS FILE IS A TEST OF HTTP GET!
			THIS IS NOT TO BE USED IN ANY PRODUCTION ENVIROMENT!
*/

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//ProsjecanTecajHrkUsd koristi REST API sa adrese  https://exchangeratesapi.io te vraća prosječan tečaj HRK u odnosu na USD za zadnjih 5 radnih dana.
func ProsjecanTecajHrkUsd() float64 {
	return GetRatesAveraged("HRK", "USD", 5)
}

//GetRatesAveraged is a currency function,
//returns averaged last n days rate of base currency against selected currency,
//with rates as returned from https://exchangeratesapi.io
//IMPORTANT: DO NOT USE FOR ANYTHING OTHER THEN TEST OF HTTP GET!
//			there is no UTC care, no averaging overflow care, no caching, precission...
func GetRatesAveraged(baseCurrency string, currency string, numberOfDays int) float64 {
	if numberOfDays > 9 {
		log.Fatal("please, don't push this test")
	}

	dayRates, err := getRates(time.Now().AddDate(0, 0, -numberOfDays).Format("2006-01-02"), time.Now().Format("2006-01-02"), baseCurrency, currency)

	if err != nil {
		log.Fatal("problems getting remote data: ", err)
	}

	var cumulative float64
	for _, u := range dayRates {
		cumulative = cumulative + u[currency]
	}

	return cumulative / float64(numberOfDays)
}

//ReturnedExchangeratesapiJSON is returned in case we are asking history data
type ReturnedExchangeratesapiJSON struct {
	StartAt string                        `json:"start_at"`
	EndAt   string                        `json:"end_at"`
	Base    string                        `json:"base"`
	Rates   map[string]map[string]float64 `json:"rates"` //[date][value]
	Error   string                        `json:"error"` //if such
}

//daily rates as returned from https://exchangeratesapi.io as [date][value]
func getRates(startDate string, endDate string, baseCurrency string, currency string) (map[string]map[string]float64, error) {

	url := "https://api.exchangeratesapi.io/history"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("start_at", startDate)
	q.Add("end_at", endDate)
	q.Add("base", baseCurrency)
	q.Add("symbols", currency)
	req.URL.RawQuery = q.Encode()

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	returnedRates := ReturnedExchangeratesapiJSON{}

	jsonErr := json.Unmarshal(body, &returnedRates)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	if returnedRates.Error != "" {
		return returnedRates.Rates, errors.New(returnedRates.Error)
	}

	return returnedRates.Rates, nil
}
