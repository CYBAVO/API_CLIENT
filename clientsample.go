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
	exampleABIJsonToBin()
	examplePostTransaction()
}

const baseURL = "API_URL"
const myAPICode = "MY_API_CODE"
const mySecretKey = "MY_SECRET_KEY"
const sigHeader = "X-Cybavo-Apisig"

func exampleGetHealth() {
	res, err := makeRequest("GET", "/v1/eos/health", nil, nil)
	log.Println("response =", res)
	log.Println("error =", err)
}

func exampleGetBlock(res string) {
	url := fmt.Sprintf("/v1/eos/block/%s", res)

	params := []string{}
	params = append(params, fmt.Sprintf("block=%s", res))
	res, err := makeRequest("GET", url, params, nil)

	log.Println("response =", res)
	log.Println("error =", err)
}

func exampleGetAccount(account string, res string) {
	url := fmt.Sprintf("/v1/eos/account/%s/%s", account, res)

	params := []string{}
	params = append(params, fmt.Sprintf("account=%s", account))
	res, err := makeRequest("GET", url, params, nil)

	log.Println("response =", res)
	log.Println("error =", err)
}

func exampleABIJsonToBin() {
	json := `{"code":"eosio.token","action":"transfer","args":{"from":"cybovatest11","to":"cybovatest12","quantity":"1.0000 EOS","memo":"NOTE"}}`

	params := []string{"r=abi_json_to_bin"}
	res, err := makeRequest("POST", "/v1/eos/abi_json_to_bin", params, &json)

	log.Println("response =", res)
	log.Println("error =", err)
}

func examplePostTransaction() {
	transaction := `{"signatures":["SIG_K1_K4gsBzrZ5dTPrK2dv1bvwttcA7aTuFFyi4X43NDPPxExLvnDxGFpkHx8tmte22sEMKgopcBYT7dvoZgVJ7HFpyQJsrZDuo"],"compression":"none","packed_context_free_data":"","packed_trx":"897ef15ba927136993dd000000000100a6823403ea3055000000572d3ccdcd0190dd39e69a64a64100000000a8ed32322190dd39e69a64a641e05b3597d15cfd45640000000000000004454f530000000000000000000000000000000000000000000000000000000000000000000000000000"}`

	params := []string{"r=push_transaction"}
	res, err := makeRequest("POST", "/v1/eos/transaction/send", params, &transaction)

	log.Println("response =", res)
	log.Println("error =", err)
}
