package main

import (
	"fmt"
	"log"
)

func main() {
	exampleGetHealth()
	exampleGetBlock("2354428")
	exampleGetBlock("latest")
	exampleGetBlock("latest_irreversible")
	exampleGetAccount("cybavotest12", "info")
	exampleGetAccount("cybavotest12", "resource")
	exampleGetAccount("cybavotest12", "balance")
}

const baseURL = "API_URL"
const myAPICode = "MY_API_CODE"
const mySecretKey = "MY_SECRET_KEY"
const apiHeader = "X-Cybavo-Apicode"
const sigHeader = "X-Cybavo-Apisig"

func exampleGetHealth() {
	res, err := makeRequest("GET", "/v1/eos/health", nil)
	log.Println("response =", res)
	log.Println("error =", err)
}

func exampleGetBlock(res string) {
	url := fmt.Sprintf("/v1/eos/block/%s", res)

	params := []string{}
	params = append(params, fmt.Sprintf("block=%s", res))
	res, err := makeRequest("GET", url, params)

	log.Println("response =", res)
	log.Println("error =", err)
}

func exampleGetAccount(account string, res string) {
	url := fmt.Sprintf("/v1/eos/account/%s/%s", account, res)

	params := []string{}
	params = append(params, fmt.Sprintf("account=%s", account))
	res, err := makeRequest("GET", url, params)

	log.Println("response =", res)
	log.Println("error =", err)
}
