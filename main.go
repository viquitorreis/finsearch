package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CompanyProfile struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Ceo               string  `json:"ceo"`
	Industry          string  `json:"industry"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
}

type CompanyMktCap struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"companyName"`
	MarketCap   float64 `json:"marketCap"`
	Sector      string  `json:"sector"`
}

type BrapiTicker struct {
	Results []struct {
		Symbol             string  `json:"symbol"`
		MarketCap          float64 `json:"marketCap"`
		ShortName          string  `json:"shortName"`
		LongName           string  `json:"longName"`
		RegularMarketPrice float64 `json:"regularMarketPrice"`
	}
}

func main() {
	// apiKey := "3sRm5UKmkVsnC0fTvlASpoueSIZCzQTv"
	// reqUrl := fmt.Sprintf("https://financialmodelingprep.com/api/v3/profile/AGRO3?apikey=%s", apiKey)
	// resp, err := http.Get(reqUrl)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// var profile []CompanyProfile
	// err = json.NewDecoder(resp.Body).Decode(&profile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	// if len(profile) > 0 {
	// 	company := profile[0]
	// 	fmt.Printf("Symbol: %s\n", company.Symbol)
	// 	fmt.Printf("Price: %.2f\n", company.Price)
	// 	fmt.Printf("CEO: %s\n", company.Ceo)
	// 	fmt.Printf("Industry: %s\n", company.Industry)
	// 	fmt.Printf("Full Time Employees: %s\n", company.FullTimeEmployees)
	// }

	// mktCap := fmt.Sprintf("https://financialmodelingprep.com/api/v3/stock-screener?apikey=%s", apiKey)
	// resp, err = http.Get(mktCap)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// var mktCapResp []CompanyMktCap
	// err = json.NewDecoder(resp.Body).Decode(&mktCapResp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, company := range mktCapResp {
	// 	fmt.Printf("Symbol: %s\n", company.Symbol)
	// 	fmt.Printf("Company Name: %s\n", company.CompanyName)
	// 	fmt.Printf("Market Cap: %.2f\n", company.MarketCap)
	// 	fmt.Printf("Sector: %s\n", company.Sector)
	// 	fmt.Println("--------------------------------------------------")
	// }
	brapiKey := "1p3Kt3LRrd6u4R9KzmuKPb"
	company := "AGRO3"
	brapiTicker := fmt.Sprintf("https://brapi.dev/api/quote/%s?token=%s", company, brapiKey)
	resp, err := http.Get(brapiTicker)
	if err != nil {
		fmt.Println(err)
		return
	}
	var tickerValues BrapiTicker
	rawBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(rawBody))

	resp.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	err = json.NewDecoder(resp.Body).Decode(&tickerValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(tickerValues.Results) > 0 {
		result := tickerValues.Results[0]
		fmt.Printf("Symbol: %s\n", result.Symbol)
		fmt.Printf("Price: %.2f\n", result.RegularMarketPrice)
		fmt.Printf("Company Name: %s\n", result.LongName)
		fmt.Printf("Market Cap: %.2f\n", result.MarketCap)
		fmt.Println("-------------------")
	} else {
		fmt.Println("Nenhum resultado encontrado")
	}
}
