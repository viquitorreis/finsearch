package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CompanyProfile struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Ceo               string  `json:"ceo"`
	Industry          string  `json:"industry"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
}

func main() {
	apiKey := "3sRm5UKmkVsnC0fTvlASpoueSIZCzQTv"
	reqUrl := fmt.Sprintf("https://financialmodelingprep.com/api/v3/profile/MSFT?apikey=%s", apiKey)
	resp, err := http.Get(reqUrl)
	// resp, err := http.Get("https://financialmodelingprep.com/api/v3/profile/AAPL?apikey=3sRm5UKmkVsnC0fTvlASpoueSIZCzQTv")
	if err != nil {
		fmt.Println(err)
		return
	}
	var profile []CompanyProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if len(profile) > 0 {
		company := profile[0]
		fmt.Printf("Symbol: %s\n", company.Symbol)
		fmt.Printf("Price: %.2f\n", company.Price)
		fmt.Printf("CEO: %s\n", company.Ceo)
		fmt.Printf("Industry: %s\n", company.Industry)
		fmt.Printf("Full Time Employees: %s\n", company.FullTimeEmployees)
	}

	// fmt.Println(jsonResp)
}
