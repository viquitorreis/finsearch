package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Iuptec/tupa"
	"github.com/joho/godotenv"
)

func HandleListAll(tc *tupa.TupaContext) error {
	godotenv.Load()
	brapiKey := os.Getenv("BRAPI_KEY")
	brapiList := "https://brapi.dev/api/quote/list"

	client := &http.Client{}
	req, err := http.NewRequest("GET", brapiList, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+brapiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var response Response
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	(*tc.Response()).Header().Set("Content-Type", "application/json")

	// Encode do response como JSON e escrevendo isso no response body
	err = json.NewEncoder(*tc.Response()).Encode(response)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func HandleTicker(tc *tupa.TupaContext) error {
	// https://brapi.dev/api/available
	godotenv.Load()
	brapiKey := os.Getenv("BRAPI_KEY")
	ticker := tc.Param("ticker")
	brapiList := fmt.Sprintf("https://brapi.dev/api/quote/%s", ticker)

	client := &http.Client{}
	req, err := http.NewRequest("GET", brapiList, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+brapiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	(*tc.Response()).Header().Set("Content-Type", "application/json")

	_, err = (*tc.Response()).Write(bodyBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
