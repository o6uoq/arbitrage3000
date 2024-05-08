package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/muesli/termenv"
)

type ExchangeRates struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func main() {
	token := strings.Trim(os.Getenv("OPENEXCHANGERATES_TOKEN"), "\"")
	if token == "" {
		fmt.Println("OPENEXCHANGERATES_TOKEN is not set")
		return
	}

	exchangeRates := fetchExchangeRates(token)
	if exchangeRates == nil || len(exchangeRates.Rates) == 0 {
		fmt.Println("Failed to fetch or parse exchange rates, or no rates data received")
		return
	}

	// ProTip: Gold in the UK is not subject to 20% VAT.
	//         Therefore, Gold is cheaper to buy in the UK.
	//         Silver is subject to VAT, therefore Silver is cheaper to buy in the US.
	proTip := `ProTip: Gold in the UK is not subject to 20% VAT.
             Therefore, Gold is cheaper to buy in the UK.
             Silver is subject to VAT, therefore Silver is cheaper to buy in the US.`
	fmt.Println(termenv.String(proTip).Foreground(termenv.ANSIGreen))

	fmt.Println() // New line after the prompt

	colorProfile := termenv.ColorProfile()
	headerColor := colorProfile.Color("5")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{termenv.String("Currency Pair").Foreground(headerColor), "Exchange Rate", "Arbitrage"})
	for _, pair := range []string{"GBP/USD", "USD/GBP", "USD/BTC", "GBP/BTC", "GBP/XAU", "GBP/XAG", "USD/XAU", "USD/XAG"} {
		rate := calculateRate(pair, exchangeRates.Rates)
		t.AppendRow(table.Row{pair, fmt.Sprintf("%.6f", rate), ""})
	}
	t.Render()

	fmt.Println() // New line between the ProTip and the first table

	t2 := table.NewWriter()
	t2.SetOutputMirror(os.Stdout)
	t2.SetStyle(table.StyleColoredBright)
	t2.AppendHeader(table.Row{"Commodity", "GBP (Sharps Pixley)", "USD (GoldSilver)"})
	t2.AppendRow(table.Row{"1oz Gold", fmt.Sprintf("£%.2f", fetchGoldPrice("https://www.sharpspixley.com/buy-bullion/buy-gold/gold-bars/1-oz-gold-bar-sharps-pixley", `£<span data="1">(.+?)<\/span>`)), fmt.Sprintf("$%.2f", fetchGoldPrice("https://goldsilver.com/buy-online/1-oz-gold-bar/", `<span class="span-big pull-right">\$(.+?)<\/span>`))})
	t2.Render()
}

func fetchExchangeRates(token string) *ExchangeRates {
	url := "https://openexchangerates.org/api/latest.json?symbols=GBP,USD,BTC,XAU,XAG"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching exchange rates:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var rates ExchangeRates
	if err := json.Unmarshal(body, &rates); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return &rates
}

func calculateRate(pair string, rates map[string]float64) float64 {
	parts := strings.Split(pair, "/")
	currency1, currency2 := parts[0], parts[1]

	if rate, found := rates[currency2]; found && currency1 == "USD" {
		return rate
	} else if rate, found := rates[currency1]; found && currency2 == "USD" {
		return 1 / rate
	} else if rate1, found1 := rates[currency1]; found1 {
		if rate2, found2 := rates[currency2]; found2 {
			return rate1 / rate2 // Cross-rate calculation, assuming rates are based on USD
		}
	}
	return 0 // Handle any cases where rates are not found
}

func fetchGoldPrice(url, regex string) float64 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching gold price:", err)
		return 0.0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0.0
	}

	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		fmt.Println("Price substring not found")
		return 0.0
	}

	price, err := strconv.ParseFloat(strings.ReplaceAll(matches[1], ",", ""), 64)
	if err != nil {
		fmt.Println("Error parsing gold price:", err)
		return 0.0
	}

	return price
}

