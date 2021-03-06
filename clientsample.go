// Copyright (c) 2018-2019 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// EOS
	exampleGetHealth()
	exampleGetBlock()
	exampleGetBlockLatest()
	exampleGetBlockLatestIrreversible()
	exampleGetAccount()
	exampleGetCurrencyBalance()
	exampleABIJsonToBin()
	examplePostPackedTransaction()
	examplePostRawTransaction()
	exampleGetInfo()

	// TRON
	exampleTronGetAccountNet()
	exampleTronGetAccount()
	exampleTronGetBlockByNum()
	exampleTronGetBlockByID()
	exampleTronGetBlockByLatestNum()
	exampleTronGetBlockByLimitNext()
	exampleTronGetNowBlock()
	exampleTronBroadcastTransaction()
	exampleTronValidateAddress()
	exampleTronGettransactionInfoByID()
	exampleTronGetTransactionByID()
	exampleTronGetContract()
	exampleTronTriggerSmartContact()
	exampleTronCreateTransaction()
}

const baseURL = "API_URL"
const myAPICode = "MY_API_CODE"
const mySecretKey = "MY_SECRET_KEY"
const sigHeader = "X-Cybavo-Apisig"

// EOS
func exampleGetHealth() {
	res, err := makeRequest("GET", "/v1/eos/health", nil, nil)
	log.Println("response =", string(res))
	log.Println("error =", err)
}

func exampleGetBlock() {
	json := `{"block_num_or_id":"2354428"}`
	resp, err := makeRequest("POST", "/v1/eos/chain/get_block", nil, &json)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetBlockLatest() {
	json := `{"block_num_or_id":"latest"}`
	resp, err := makeRequest("POST", "/v1/eos/chain/get_block", nil, &json)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetBlockLatestIrreversible() {
	json := `{"block_num_or_id":"latest_irreversible"}`
	resp, err := makeRequest("POST", "/v1/eos/chain/get_block", nil, &json)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetAccount() {
	json := `{"account_name":"cybavotest12"}`
	resp, err := makeRequest("POST", "/v1/eos/chain/get_account", nil, &json)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetCurrencyBalance() {
	json := `{"code":"eosio.token","account":"cybavotest12","symbol":"EOS"}`
	resp, err := makeRequest("POST", "/v1/eos/chain/get_currency_balance", nil, &json)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleABIJsonToBin() {
	json := `{"code":"eosio.token","action":"transfer","args":{"from":"cybovatest11","to":"cybovatest12","quantity":"1.0000 EOS","memo":"NOTE"}}`

	res, err := makeRequest("POST", "/v1/eos/chain/abi_json_to_bin", nil, &json)

	log.Println("response =", string(res))
	log.Println("error =", err)
}

func examplePostPackedTransaction() {
	transaction := `{"signatures":["SIG_K1_K4gsBzrZ5dTPrK2dv1bvwttcA7aTuFFyi4X43NDPPxExLvnDxGFpkHx8tmte22sEMKgopcBYT7dvoZgVJ7HFpyQJsrZDuo"],"compression":"none","packed_context_free_data":"","packed_trx":"897ef15ba927136993dd000000000100a6823403ea3055000000572d3ccdcd0190dd39e69a64a64100000000a8ed32322190dd39e69a64a641e05b3597d15cfd45640000000000000004454f530000000000000000000000000000000000000000000000000000000000000000000000000000"}`

	params := []string{"fmt=packed"}
	res, err := makeRequest("POST", "/v1/eos/chain/push_transaction", params, &transaction)

	log.Println("response =", string(res))
	log.Println("error =", err)
}

func examplePostRawTransaction() {
	transaction := `{"expiration":"2019-04-17T00:00:00","ref_block_num":64484,"ref_block_prefix":2668208063,"max_net_usage_words":0,"max_cpu_usage_ms":0,"delay_sec":0,"context_free_actions":[],"actions":[{"account":"eosio.token","name":"transfer","authorization":[{"actor":"cybovatest11","permission":"active"}],"data":"00afa998aaabaac6a0986aff4b9a3c61102700000000000004454f5300000000057465737431"}],"transaction_extensions":[],"signatures":["SIG_K1_K3p94niNvkxzpMYezEzetcoFTEyowgVaX95p5K8xqEdyP2pFkcvqeVXbyMZMBWBDe73G5Dv92SLyTBxaj5yNnStALET326"],"context_free_data":[]}`

	params := []string{"fmt=raw"}
	res, err := makeRequest("POST", "/v1/eos/chain/push_transaction", params, &transaction)

	log.Println("response =", string(res))
	log.Println("error =", err)
}

func exampleGetInfo() {
	resp, err := makeRequest("POST", "/v1/eos/chain/get_info", nil, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

// TRON
type response struct {
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Result       string `json:"result,omitempty"`
}

func printResult(result []byte, err error) {
	if err == nil {
		var res response
		_ = json.Unmarshal(result, &res)
		if res.ErrorCode != 0 {
			log.Println(res.ErrorMessage)
		} else {
			log.Println(res.Result)
		}
	} else {
		log.Println("error =", err)
	}
}

func exampleTronGetAccountNet() {
	request := `{"address": "4193A8BC2E7D6BB1BD75FB2D74107FFBDA81AF439D"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getaccountnet", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetAccount() {
	request := `{"address": "4193A8BC2E7D6BB1BD75FB2D74107FFBDA81AF439D"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getaccount", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetBlockByNum() {
	request := `{"num":200}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getblockbynum", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetBlockByID() {
	request := `{"value": "00000000000000c83ad3d3759a57db65c945053a0351bd044e3cb46dd939ecfd"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getblockbyid", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetBlockByLatestNum() {
	request := `{"num":1}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getblockbylatestnum", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetBlockByLimitNext() {
	request := `{"startNum":1,"endNum":2}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getblockbylimitnext", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetNowBlock() {
	request := `{"req":"nowblock"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getnowblock", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronBroadcastTransaction() {
	request := `{"raw_data":{"ref_block_bytes":"e6dc","ref_block_hash":"cb95b9e4914854cc","expiration":1552908215000,"timestamp":1552907315000,"contract":[{"type":"TransferContract","parameter":{"value":{"amount":50000000,"owner_address":"412a00fc3f6068ed005d3d00b30609ddb79c26e213","to_address":"412be614e143e40979efd18b47cfe089c2ba5906f1"},"type_url":"type.googleapis.com/protocol.TransferContract"}}]},"signature":["edaa7bd24c09431ff423194f60730e421e5e207f747a8c1c738e370d2b779bf12677a3dc34b98afb6bc997d0b53175a4ac945dd7a571ec07138d243532840d4200"],"txID":"1d3e4cfa27460a0b0ba3e3a8eaae61b8786979fef2cdbeea0df550fb60369e9d"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/broadcasttransaction", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronValidateAddress() {
	request := `{"address": "4170622A5F83CF2A5DE36433AD4F6CDEF078B01F05"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/validateaddress", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGettransactionInfoByID() {
	request := `{"value":"9af66ca310ffb0f450c36128be690a65dc947eaeef424bb1512ad7b4e5958044"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/gettransactioninfobyid", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetTransactionByID() {
	request := `{"value": "9af66ca310ffb0f450c36128be690a65dc947eaeef424bb1512ad7b4e5958044"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/gettransactionbyid", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetContract() {
	request := `{"value":"41909C9BFAB86CBF89B1A4D1488744111637719B14"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/getcontract", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronTriggerSmartContact() {
	request := `{"contract_address":"414B9223269E1C53B6B72745E3332623C21E70B805","function_selector":"balanceOf(address)","parameter":"000000000000000000000000C12FB0805F3BA3CA3A01D258A887DD3E4DB731A6","fee_limit":1000000,"call_value":0,"owner_address":"41C12FB0805F3BA3CA3A01D258A887DD3E4DB731A6"}`

	resp, err := makeRequest("POST", "/v1/tron/wallet/triggersmartcontract", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronCreateTransaction() {
	request := `{"to_address": "412be614e143e40979efd18b47cfe089c2ba5906f1", "owner_address": "412a00fc3f6068ed005d3d00b30609ddb79c26e213", "amount": 1000 }`

	resp, err := makeRequest("POST", "/v1/tron/wallet/createtransaction", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetBlockByNum_Solidity() {
	request := `{"num":200}`

	resp, err := makeRequest("POST", "/v1/tron/walletsolidity/getblockbynum", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleTronGetNowBlock_Solidity() {
	request := `{"req":"nowblock"}`

	resp, err := makeRequest("POST", "/v1/tron/walletsolidity/getnowblock", nil, &request)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

//For LTC
func exampleLTCGetFee() {
	url := "/v1/ltc/wallet/fee"

	params := []string{}
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetBalance(address string) {
	url := fmt.Sprintf("/v1/ltc/wallet/addressbalance/%s", address)

	params := []string{}
	params = append(params, fmt.Sprintf("address=%s", address))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetConfirmBlocksByTxID(txid string) {
	url := fmt.Sprintf("/v1/ltc/transaction/confirm/%s", txid)

	params := []string{}
	params = append(params, fmt.Sprintf("txid=%s", txid))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetUnconfirmBalance(address string) {
	url := fmt.Sprintf("/v1/ltc/transaction/addressunconfirmbalance/%s", address)

	params := []string{}
	params = append(params, fmt.Sprintf("address=%s", address))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func examplePushTransaction() {
	request := `{"Type":0,"FromAddress":"QNdiW2GFR4mTw1in7Nn2LAEV1NeFphUAJg","TokenAddress":"","ToAddress":"miQihDe6EvUzaKX7DB4yFfTEB3P6FYDVWK","Amount":"0.87788769","Fee":"0.0000416"}`

	resp, err := makeRequest("POST", "/v1/ltc/transaction/createpayment", nil, &request)
	printResult(resp, err)
}

func exampleSendSignedTransaction() {
	request := `{"Type":0,"TokenAddress":"","Transaction":"01000000000101fbb28dd706c6817525a55835cbd8da78b67f617f731dbd0fbbd98460492d543d01000000171600148dc0222aad18401c9b5a510b10a3177729e6ccc0ffffffff0210270000000000001976a9141fba8742e5b47d63c100272e3d017c7972dd392688ac101bcb1d0000000017a914aed75052d9ecaf0a2562aaee8ee044dfc2bb75bd8702473044022035a599e87c928a6e27310e63c2e9e6c7b086504afb7ee44c9baa442d123375e20220023fd88234fce894484a88061cb8b213c272326c41f8d552d10055043d885476012103379b73388a9cc5a09f310d96cc357a50d1ebbe57e0ef904cf205714411eacac100000000"}`

	resp, err := makeRequest("POST", "/v1/ltc/transaction/submitpayment", nil, &request)
	printResult(resp, err)
}

func exampleGetUTXO(address string) {
	url := fmt.Sprintf("/v1/ltc/transaction/addressutxo/%s", address)

	params := []string{}
	params = append(params, fmt.Sprintf("address=%s", address))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetAddressHistory(address string, count string) {
	url := fmt.Sprintf("/v1/ltc/wallet/addresshistory/%s/%s", address, count)

	params := []string{}
	params = append(params, fmt.Sprintf("address=%s", address))
	params = append(params, fmt.Sprintf("count=%s", count))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetTxInfo(txid string, address string) {
	url := fmt.Sprintf("/v1/ltc/transaction/txinfo/%s/%s", txid, address)

	params := []string{}
	params = append(params, fmt.Sprintf("txid=%s", txid))
	params = append(params, fmt.Sprintf("address=%s", address))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleGetBatchAddressBalance() {
	request := `{"Addresses":[{"Address":"QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT","Token":""},{"Address":"QaAqKiTwm5qpYyjuSLRXhuAHtpBuWn6vFU","Token":""}]}`

	resp, err := makeRequest("POST", "/v1/ltc/wallet/addressbalance", nil, &request)
	printResult(resp, err)
}

func exampleXPubGetBalance(xpub string) {
	url := fmt.Sprintf("/v1/ltc/wallet/xpubbalance/%s", xpub)

	params := []string{}
	params = append(params, fmt.Sprintf("xpub=%s", xpub))
	resp, err := makeRequest("GET", url, params, nil)

	log.Println("response =", string(resp))
	log.Println("error =", err)
}

func exampleXPubPushTransaction() {
	request := `{"ToAddress":"QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT","Amount":"0.0001","Fee":"0.0000416","XPub":"tpubDH9DBHzyDeLhoXWfyaaWFjiDRXsYguxZHKmbqTXqRMsHyt21jA1cXhMCjuBvFo2Q6PzY4UwRS1zArQH1aaKADdASW3vyZWwvRSF9vVLvmrL","ReceivingAddress":"mkJ6dFu4pzohktR8rwH8q1eK3qs8BTsBsY"}`

	resp, err := makeRequest("POST", "/v1/ltc/transaction/xpubcreatepayment", nil, &request)
	printResult(resp, err)
}
