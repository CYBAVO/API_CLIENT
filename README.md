# CYBAVO Ledger Service
## Supportted Currencies
* [EOS](#CYBAVO-EOS-API-Sample)
  * Basic information [[go]](#Basic-information) 
  * Account information [[go]](#Account-information) 
  * Resource information [[go]](#Resource-information)
  * Balance information [[go]](#Balance-information)
  * Transactions information [[go]](#Transactions-information)
  * Transactions information in latest block [[go]](#Transactions-information-in-latest-block)
  * Transactions information in the latest irreversible block [[go]](#Transactions-information-in-the-latest-irreversible-block)
  * Convert json to binary [[go]](#Convert-json-to-binary)
  * Broadcast signed transaction [[go]](#Broadcast-signed-transaction)

* [TRON](#CYBAVO-TRON-API-Sample)
	* Get Account Information [[go]](#Get-Account-Information)
	* Account Bandwidth Information [[go]](#Account-Bandwidth-Information)
	* Block by Block Height [[go]](#Block-by-Block-Height)
	* Block by ID [[go]](#Block-by-ID)
	* Block by Latest Number [[go]](#Block-by-Latest-Number)
	* Block by Limit Next [[go]](#Block-by-Limit-Next)
	* Get Latest Block [[go]](#Get-Latest-Block)
	* Broadcast Transaction [[go]](#Broadcast-Transaction)
	* Validate Address [[go]](#Validate-Address)
	* Transactions by ID [[go]](#Transactions-by-ID)
	* Transaction Info by ID [[go]](#Transaction-Info-by-ID)
	* Get Contract [[go]](#Get-Contract)
	* Trigger Smart Contract [[go]](#Trigger-Smart-Contract)

* [Litecoin](#CYBAVO-Litecoin-API-Sample)
	* Transaction fee [[go]](#LTC---Transaction-fee)
	* Get address balance [[go]](#LTC---Get-address-balance)
	* Get confirmation of TXID [[go]](#LTC---Get-confirmation-of-TXID)
	* Get unconfirmed balance [[go]](#LTC---Get-unconfirmed-balance)
	* Create raw transaction [[go]](#LTC---Create-raw-transaction)
	* Broadcast signed transaction [[go]](#LTC---Broadcast-signed-transaction)
	* Get UTXO in address [[go]](#LTC---Broadcast-signed-transaction)
	* Get transaction history in address [[go]](#LTC---Get-transaction-history-in-address)
	* Get transaction information [[go]](#LTC---Get-transaction-information)
	* Get batch of address balance [[go]](#LTC---Get-batch-of-address-balance)
	* Get balance list of XPub [[go]](#LTC---Get-balance-list-of-XPub)
	* Create transaction blob with XPub [[go]](#LTC---Create-transaction-blob-with-XPub)

* [Bitcoin](#CYBAVO-Bitcoin-API-Sample)
	* Transaction fee [[go]](#BTC---Transaction-fee)
	* Get address balance [[go]](#BTC---Get-address-balance)
	* Get confirmation of TXID [[go]](#BTC---Get-confirmation-of-TXID)
	* Get unconfirmed balance [[go]](#BTC---Get-unconfirmed-balance)
	* Create raw transaction [[go]](#BTC---Create-raw-transaction)
	* Broadcast signed transaction [[go]](#BTC---Broadcast-signed-transaction)
	* Get UTXO in address [[go]](#BTC---Broadcast-signed-transaction)
	* Get transaction history in address [[go]](#BTC---Get-transaction-history-in-address)
	* Get transaction information [[go]](#BTC---Get-transaction-information)
	* Get block information [[go]](#BTC---Get-block-information)

## Prerequisite

- Please obtain following info from CYBAVO Team
	- API URL 
	- API code
	- API secret

---------------------------------------
# CYBAVO EOS API Sample

####  Account information
- Query basic information
	- **GET** /v1/eos/get_info
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "server_version": "2162999e",
		    "chain_id": "e70aaab8997e1dfce58fbfac80cbbb8fecec7b99cf982a9444273cbc64c41473",
		    "head_block_num": 25312519,
		    "last_irreversible_block_num": 25312188,
		    "last_irreversible_block_id": "01823bbca04d7c1eb7cb59afed326babd9815a9f2c12b02317bd8d27b237663f",
		    "head_block_id": "01823d07118e4fa3036b529cb0d170bea7cb3e3775d27ca659a6384a4179d924",
		    "head_block_time": "2019-04-24T22:47:16",
		    "head_block_producer": "ohtigertiger",
		    "virtual_block_cpu_limit": 200000000,
		    "virtual_block_net_limit": 524288000,
		    "block_cpu_limit": 199900,
		    "block_net_limit": 524288,
		    "server_version_string": "jenkins-fullnode_eos-9"
		  }
		}
		```

- Query basic information of the specific account 
	- **GET** /v1/eos/account/`EOS_ACCOUNT_NAME`/info
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "account_name": "cybavotest12",
		    "created": "2018-11-26T03:27:40",
		    "permissions": [
		      {
		        "parent": "owner",
		        "perm_name": "active",
		        "required_auth": {
		          "keys": [
		            {
		              "key": "EOS7fbBtWWoFot3Q4bBDfQ1dy8Ty6ddZBvxrPcTsuqYdrCnFRj9k2",
		              "weight": 1
		            }
		          ],
		          "threshold": 1
		        }
		      },
		      {
		        "parent": "",
		        "perm_name": "owner",
		        "required_auth": {
		          "keys": [
		            {
		              "key": "EOS7fbBtWWoFot3Q4bBDfQ1dy8Ty6ddZBvxrPcTsuqYdrCnFRj9k2",
		              "weight": 1
		            }
		          ],
		          "threshold": 1
		        }
		      }
		    ]
		  }
		}
		```

#### Resource information
- Query resource information of the specific account
	- **GET** /v1/eos/account/`EOS_ACCOUNT_NAME`/resouce
	  - Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "cpu_limit": {
		      "available": 4870,
		      "max": 4870,
		      "used": 0
		    },
		    "net_limit": {
		      "available": 27298,
		      "max": 27298,
		      "used": 0
		    },
		    "ram_quota": 5467,
		    "ram_usage": 3446,
		    "total_resources": {
		      "cpu_weight": "0.1000 EOS",
		      "net_weight": "0.1000 EOS",
		      "owner": "cybavotest12",
		      "ram_bytes": 4067
		    }
		  }
		}
		```

#### Balance information
- Query balance of the specific account
	- **GET** /v1/eos/account/`EOS_ACCOUNT_NAME`/balance
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": [
		    "2.3500 EOS"
		  ]
		}
		```

#### Transactions information
- Query detailed information of transactions in the specific block and block height
	- **GET** /v1/eos/block/`BLOCK_NUM`
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "_id": "5c6d249b9aeae4d11a59b09a",
		    "block": {
		      "action_mroot": "f4f77c2e8ac7137e94e69a4977387e6f5e7becd699b8c7af983a7094f2b0fc3e",
		      "block_extensions": [],
		      "confirmed": 0,
		      "header_extensions": [],
		      "new_producers": null,
		      "previous": "0023ecfb86f42282f800fb5fedfedc7e6756b863c4878c4a93e2118cfb23f09a",
		      "producer": "junglemorpho",
		      "producer_signature": "SIG_K1_K3Fantwu8VHApcUhR2M64fMa11WotLFK5Eh3JEEHrkN9bY6hpYcdQkZpfUkU6qoMbZaHV2EQ1BWZxox3DyKU8ibEnZLJSs",
		      "schedule_version": 54,
		      "timestamp": "2018-12-07T11:54:58.500",
		      "transaction_mroot": "07cc36a81694a441e0bc891b9d25c59369a4929e3522b16d5283f6e59f68fecc",
		      "transactions": [
		        {
		          "cpu_usage_us": 485,
		          "net_usage_words": 16,
		          "status": "executed",
		          "trx": {
		            "compression": "none",
		            "context_free_data": [],
		            "id": "af6bc4ce0992f9ab467d55650112b9d75648fc1c9e010b0ad88f92f05153b2e4",
		            "packed_context_free_data": "",
		            "packed_trx": "af5f0a5cfaec9b6596ec000000000100a6823403ea3055000000572d3ccdcd011042c62ad36d8e4700000000a8ed3232211042c62ad36d8e472042c62ad36d8e47e80300000000000004454f53000000000000",
		            "signatures": [
		              "SIG_K1_KBY7pdHXHrq1cRLgFyqvw9MRavqGRNqqFbfyKzfmmPacW2difPPs8LiejJqBLWuhkMAQg8867QUpPmE3YMJLPczV5jtYDa"
		            ],
		            "transaction": {
		              "actions": [
		                {
		                  "account": "eosio.token",
		                  "authorization": [
		                    {
		                      "actor": "cybavotest11",
		                      "permission": "active"
		                    }
		                  ],
		                  "data": {
		                    "from": "cybavotest11",
		                    "memo": "",
		                    "quantity": "0.1000 EOS",
		                    "to": "cybavotest12"
		                  },
		                  "hex_data": "1042c62ad36d8e472042c62ad36d8e47e80300000000000004454f530000000000",
		                  "name": "transfer"
		                }
		              ],
		              "context_free_actions": [],
		              "delay_sec": 0,
		              "expiration": "2018-12-07T11:55:27",
		              "max_cpu_usage_ms": 0,
		              "max_net_usage_words": 0,
		              "ref_block_num": 60666,
		              "ref_block_prefix": 3969279387,
		              "transaction_extensions": []
		            }
		          }
		        }
		      ]
		    },
		    "block_id": "0023ecfc3a2d3ad72e8d159a90d176db5e0ff55e6ede41ed9a687637d2c9d1a9",
		    "block_num": 2354428,
		    "createdAt": "2019-02-20T09:57:47.569Z",
		    "irreversible": true,
		    "updatedAt": "2019-02-20T09:57:49.654Z",
		    "validated": true
		  }
		}
		```

#### Transactions information in the latest block
- Query detailed information of transactions in the latest block and block height
	- **GET** /v1/eos/block/latest
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "timestamp": "2019-03-04T05:42:18",
		    "producer": "eosiodetroit",
		    "confirmed": 226,
		    "previous": "00fe66d354f86c96e53da8eab37e2ca4b5e3d172252872fc45f68c83f748e9aa",
		    "transaction_mroot": "c8449ad621871aba6b2f52e7c0e5437b104df01f1db545f241fbc53cb2663b79",
		    "action_mroot": "56d1c5520bdd326cf7c34dee0956749defe91cf968afe5d03d99237e16def047",
		    "schedule_version": 119,
		    "new_producers": null,
		    "header_extensions": [],
		    "producer_signature": "SIG_K1_KBoy7P4uWwk4vqYzLP3qp9rmNCfgdBxw9TWGAqUHmVjTzmCvmaTxKFj917yMCTWpJRGw8z1HL7VCkucLYxMQ4GvAcys943",
		    "transactions": [
		      {
			"status": "executed",
			"cpu_usage_us": 1470,
			"net_usage_words": 27,
			"trx": [
			  "fe91ac50f28622eb8ba4fa8d43c5998822943f3da4d00b621a00ffe0507682f2",
			  {
			    "signatures": [
			      "SIG_K1_K4RY1aNGHNvFw3gJCJGJB2trsy7byD1ojiia44Q1pgoK6DbMdVmWxwpqfbHanDrGdwTEhCjN93nTDUvVzTpLsRsgJSFhU7"
			    ],
			    "compression": "none",
			    "packed_context_free_data": "",
			    "packed_trx": "f4ba7c5c8765a1f04907000000000210b2b9062173b2390000c8eaa8ed3232011042a6875e7a90c300000000a8ed3232331cad0100000000001042a6875e7a90c3e80300000000000004454f5300000000010000e80300000000000004454f530000000000a6823403ea3055000000572d3ccdcd011042a6875e7a90c300000000a8ed3232211042a6875e7a90c310b2b9062173b239e80300000000000004454f53000000000000"
			  }
			]
		      },
		      {
			"status": "executed",
			"cpu_usage_us": 1323,
			"net_usage_words": 27,
			"trx": [
			  "a3163d4db75704e49209202a00d18f3bc9b91efb2dc6534d618c2d9734408265",
			  {
			    "signatures": [
			      "SIG_K1_KACxaNXx7L78Hcp93sUNMr6iNsMyaxK6jMyYh72r4U1XnFcXyw1s3vbrx2LEiwLut3NLUNrpqQ4tcZBsmgcJbu37wUppoJ"
			    ],
			    "compression": "none",
			    "packed_context_free_data": "",
			    "packed_trx": "f4ba7c5c8765a1f04907000000000210b2b9062173b2390000c8eaa8ed3232012042a6875e7a90c300000000a8ed3232331cad0100000000002042a6875e7a90c3e80300000000000004454f5300000000010100e80300000000000004454f530000000000a6823403ea3055000000572d3ccdcd012042a6875e7a90c300000000a8ed3232212042a6875e7a90c310b2b9062173b239e80300000000000004454f53000000000000"
			  }
			]
		      },
		      {
			"status": "executed",
			"cpu_usage_us": 1378,
			"net_usage_words": 27,
			"trx": [
			  "2bbeee7559d916f361aba1ddeb8ce6ec852d441001488ef706e5afd734da0209",
			  {
			    "signatures": [
			      "SIG_K1_KmRMShYWbo6SHp1c56m8zHvmFaYvrqHwSitKV25ovpbiQXwP4uWMQZzEhuQMF6X24HJNqKMtXdWRK87Pbd13xnWAbr8NPL"
			    ],
			    "compression": "none",
			    "packed_context_free_data": "",
			    "packed_trx": "f4ba7c5c8765a1f04907000000000210b2b9062173b2390000c8eaa8ed3232015042a6875e7a90c300000000a8ed3232331cad0100000000005042a6875e7a90c3e80300000000000004454f5300000000010000e80300000000000004454f530000000000a6823403ea3055000000572d3ccdcd015042a6875e7a90c300000000a8ed3232215042a6875e7a90c310b2b9062173b239e80300000000000004454f53000000000000"
			  }
			]
		      },
		      {
			"status": "executed",
			"cpu_usage_us": 1483,
			"net_usage_words": 27,
			"trx": [
			  "d98b49a280512d7cdcd3ae87463210be5a13ed6201c901fa175fb7b15787d37b",
			  {
			    "signatures": [
			      "SIG_K1_K5dCvGNJVA54kah8SiRfk4rw8s5S3jEFii93sB7WVg2jzTaUwHa6eXaHgBrMPEoDTBy3ekTXdA2LMH6dA1WfGaCbwqLM9v"
			    ],
			    "compression": "none",
			    "packed_context_free_data": "",
			    "packed_trx": "f4ba7c5c8765a1f04907000000000210b2b9062173b2390000c8eaa8ed3232013042a6875e7a90c300000000a8ed3232331cad0100000000003042a6875e7a90c3e80300000000000004454f5300000000010000e80300000000000004454f530000000000a6823403ea3055000000572d3ccdcd013042a6875e7a90c300000000a8ed3232213042a6875e7a90c310b2b9062173b239e80300000000000004454f53000000000000"
			  }
			]
		      },
		      {
			"status": "executed",
			"cpu_usage_us": 1526,
			"net_usage_words": 27,
			"trx": [
			  "02896e12d04a6bc1d310ce182d4c7aa1ece3e74872a102d514d340c7a37f5a92",
			  {
			    "signatures": [
			      "SIG_K1_KisNJbXu5Be7J3fVRvx9pG9eRdnkWaFTn6jickJFfgPtcWvDCDr3A17LrVYv4JUAb4in4pxvdiJsNNJB4475vYtio7DXuc"
			    ],
			    "compression": "none",
			    "packed_context_free_data": "",
			    "packed_trx": "f4ba7c5c8765a1f04907000000000210b2b9062173b2390000c8eaa8ed3232014042a6875e7a90c300000000a8ed3232331cad0100000000004042a6875e7a90c3e80300000000000004454f5300000000010100e80300000000000004454f530000000000a6823403ea3055000000572d3ccdcd014042a6875e7a90c300000000a8ed3232214042a6875e7a90c310b2b9062173b239e80300000000000004454f53000000000000"
			  }
			]
		      },
		      {
			"status": "executed",
			"cpu_usage_us": 24553,
			"net_usage_words": 28,
			"trx": [
			  "a4cfcd977d930aa65465932dc9aa49839feb797e0d6436eae91437c910e801f1",
			  null
			]
		      },
		      {
			"status": "soft_fail",
			"cpu_usage_us": 17857,
			"net_usage_words": 0,
			"trx": [
			  "d84727d4b20e0bab2b4e460d83d5204650d7c810606f0f1b94eef885459dfe6d",
			  null
			]
		      },
		      {
			"status": "soft_fail",
			"cpu_usage_us": 17814,
			"net_usage_words": 0,
			"trx": [
			  "6ba29f1b6b4fd542f46990a3870bb1e359478d913c05886e24ab9e114e35ef81",
			  null
			]
		      },
		      {
			"status": "soft_fail",
			"cpu_usage_us": 17759,
			"net_usage_words": 0,
			"trx": [
			  "7809df7480cc6b8d411258c9e58c539dae93f5b17f1363f21e6b7c1a3e09a694",
			  null
			]
		      },
		      {
			"status": "soft_fail",
			"cpu_usage_us": 17852,
			"net_usage_words": 0,
			"trx": [
			  "dba332bb26852dd7c638a8c733b770bc4a7b13fc9e12639c8922939bc134d709",
			  null
			]
		      }
		    ],
		    "id": "00fe66d4b654ce40de0178e0df26b34a9bd04d8fbe49bf34ed0c1ec10fb4ea24",
		    "block_num": 16672468,
		    "ref_block_prefix": 3765961182,
		    "block_extensions": []
		  }
		}
		```

#### Transactions information in the latest irreversible block
- Query detailed information of transactions in the latest irreversible block and block height
	- **GET** /v1/eos/block/latest\_irreversible
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "timestamp": "2019-03-04T05:39:29",
		    "producer": "atticlabjbpn",
		    "confirmed": 0,
		    "previous": "00fe6592e1fe03eb8e6123a82fb4564f227d2623845140fb0952e1c83201f673",
		    "transaction_mroot": "0000000000000000000000000000000000000000000000000000000000000000",
		    "action_mroot": "ffefc5d30aa64318db8ad0873d0d7ab3ca984c424bbb992716350171ef92aff5",
		    "schedule_version": 119,
		    "new_producers": null,
		    "header_extensions": [],
		    "producer_signature": "SIG_K1_KchKpD4g5hjRP9KcdPYAXkKwSzRvoYUFdUcMmTFTAV5REAZrQAQaZQ6VuktEeo4w4ZPWVCQyTji1sTzUdJQGTUtzRyFejB",
		    "transactions": [],
		    "id": "00fe65938f82710c6ca3f180eb674ed2e8df933e58babb1bb01b1785a59eebb9",
		    "block_num": 16672147,
		    "ref_block_prefix": 2163319660,
		    "block_extensions": []
		  }
		}
		```

#### Convert json to binary
- Convert json to binary
	- **POST** /v1/eos/abi\_json\_to\_bin
		- Request
		
		``` json
		{
		  "code": "eosio.token",
		  "action": "transfer",
		  "args": {
		    "from": "cybovatest11",
		    "to": "cybovatest12",
		    "quantity": "1.0000 EOS",
		    "memo": "NOTE"
		  }
		}
		```

		- Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "binargs": "1042c62a9b4d8f472042c62a9b4d8f47102700000000000004454f5300000000044e4f5445"
		  }
		}
		```

#### Broadcast signed transaction
- Receive the signed JSON transaction and broadcast it to blockchain
	- **POST** /v1/eos/transaction/send?fmt=packed
		- Request

		``` json
		{
		  "signatures": [
		    "SIG_K1_K3p94niNvkxzpMYezEzetcoFTEyowgVaX95p5K8xqEdyP2pFkcvqeVXbyMZMBWBDe73G5Dv92SLyTBxaj5yNnStALET326"
		  ],
		  "compression": "none",
		  "packed_context_free_data": "",
		  "packed_trx": "3a2a825c437e17a06a68000000000100a6823403ea3055000000572d3ccdcd011042421a7b53315500000000a8ed3232211042421a7b5331551042746679533155e80300000000000004454f53000000000000"
		}
		```
		
		-	Response

		``` json
		{
		  "error_code": 0,
		  "result": {
		    "transaction_id": "d9837444fc972a0a8af386170a83796d9c064705bf8ae51e5272768581f52822",
		    "processed": {
		      "id": "d9837444fc972a0a8af386170a83796d9c064705bf8ae51e5272768581f52822",
		      "block_num": 17334041,
		      "block_time": "2019-03-08T08:36:21.500",
		      "producer_block_id": null,
		      "receipt": {
		        "status": "executed",
		        "cpu_usage_us": 624,
		        "net_usage_words": 16
		      },
		      "elapsed": 624,
		      "net_usage": 128,
		      "scheduled": false,
		      "action_traces": [
		        {
		          "receipt": {
		            "receiver": "eosio.token",
		            "act_digest": "f88c491d8440b65532a1d060c417668cdfe3f32cf6d77ccada0f240c1ab20c14",
		            "global_sequence": 227073094,
		            "recv_sequence": 41071850,
		            "auth_sequence": [
		              [
		                "eospaysucd11",
		                10
		              ]
		            ],
		            "code_sequence": 3,
		            "abi_sequence": 2
		          },
		          "act": {
		            "account": "eosio.token",
		            "name": "transfer",
		            "authorization": [
		              {
		                "actor": "eospaysucd11",
		                "permission": "active"
		              }
		            ],
		            "data": {
		              "from": "eospaysucd11",
		              "to": "eospayfail11",
		              "quantity": "0.1000 EOS",
		              "memo": ""
		            },
		            "hex_data": "1042421a7b5331551042746679533155e80300000000000004454f530000000000"
		          },
		          "context_free": false,
		          "elapsed": 336,
		          "console": "",
		          "trx_id": "d9837444fc972a0a8af386170a83796d9c064705bf8ae51e5272768581f52822",
		          "block_num": 17334041,
		          "block_time": "2019-03-08T08:36:21.500",
		          "producer_block_id": null,
		          "account_ram_deltas": [],
		          "except": null,
		          "inline_traces": [
		            {
		              "receipt": {
		                "receiver": "eospaysucd11",
		                "act_digest": "f88c491d8440b65532a1d060c417668cdfe3f32cf6d77ccada0f240c1ab20c14",
		                "global_sequence": 227073095,
		                "recv_sequence": 6,
		                "auth_sequence": [
		                  [
		                    "eospaysucd11",
		                    11
		                  ]
		                ],
		                "code_sequence": 3,
		                "abi_sequence": 2
		              },
		              "act": {
		                "account": "eosio.token",
		                "name": "transfer",
		                "authorization": [
		                  {
		                    "actor": "eospaysucd11",
		                    "permission": "active"
		                  }
		                ],
		                "data": {
		                  "from": "eospaysucd11",
		                  "to": "eospayfail11",
		                  "quantity": "0.1000 EOS",
		                  "memo": ""
		                },
		                "hex_data": "1042421a7b5331551042746679533155e80300000000000004454f530000000000"
		              },
		              "context_free": false,
		              "elapsed": 15,
		              "console": "",
		              "trx_id": "d9837444fc972a0a8af386170a83796d9c064705bf8ae51e5272768581f52822",
		              "block_num": 17334041,
		              "block_time": "2019-03-08T08:36:21.500",
		              "producer_block_id": null,
		              "account_ram_deltas": [],
		              "except": null,
		              "inline_traces": []
		            },
		            {
		              "receipt": {
		                "receiver": "eospayfail11",
		                "act_digest": "f88c491d8440b65532a1d060c417668cdfe3f32cf6d77ccada0f240c1ab20c14",
		                "global_sequence": 227073096,
		                "recv_sequence": 4,
		                "auth_sequence": [
		                  [
		                    "eospaysucd11",
		                    12
		                  ]
		                ],
		                "code_sequence": 3,
		                "abi_sequence": 2
		              },
		              "act": {
		                "account": "eosio.token",
		                "name": "transfer",
		                "authorization": [
		                  {
		                    "actor": "eospaysucd11",
		                    "permission": "active"
		                  }
		                ],
		                "data": {
		                  "from": "eospaysucd11",
		                  "to": "eospayfail11",
		                  "quantity": "0.1000 EOS",
		                  "memo": ""
		                },
		                "hex_data": "1042421a7b5331551042746679533155e80300000000000004454f530000000000"
		              },
		              "context_free": false,
		              "elapsed": 23,
		              "console": "",
		              "trx_id": "d9837444fc972a0a8af386170a83796d9c064705bf8ae51e5272768581f52822",
		              "block_num": 17334041,
		              "block_time": "2019-03-08T08:36:21.500",
		              "producer_block_id": null,
		              "account_ram_deltas": [],
		              "except": null,
		              "inline_traces": []
		            }
		          ]
		        }
		      ],
		      "except": null
		    }
		  }
		}
		```
	- **POST** /v1/eos/transaction/send?fmt=raw
		- Request

		``` json
		{
		  "expiration": "2019-04-17T00:00:00",
		  "ref_block_num": 64484,
		  "ref_block_prefix": 2668208063,
		  "max_net_usage_words": 0,
		  "max_cpu_usage_ms": 0,
		  "delay_sec": 0,
		  "context_free_actions": [],
		  "actions": [
		    {
		      "account": "eosio.token",
		      "name": "transfer",
		      "authorization": [
		        {
		          "actor": "cybovatest11",
		          "permission": "active"
		        }
		      ],
		      "data": "00afa998aaabaac6a0986aff4b9a3c61102700000000000004454f5300000000057465737431"
		    }
		  ],
		  "transaction_extensions": [],
		  "signatures": [
		    "SIG_K1_K3p94niNvkxzpMYezEzetcoFTEyowgVaX95p5K8xqEdyP2pFkcvqeVXbyMZMBWBDe73G5Dv92SLyTBxaj5yNnStALET326"
		  ],
		  "context_free_data": []
		}
		```

---------------------------------------
# CYBAVO TRON API Sample

### Get Account Information
- Queries information about an account. Returns the account object.
	- **POST** /v1/tron/wallet/getaccount
		-  Request

		``` json
		{
		  "address": "4170622A5F83CF2A5DE36433AD4F6CDEF078B01F05"
		}
		```
		
		-	Response

		``` json
		{
		  "address": "4170622a5f83cf2a5de36433ad4f6cdef078b01f05",
		  "balance": 166600000,
		  "create_time": 1551773187000,
		  "account_resource": {},
		  "assetV2": [
		    {
		      "key": "1000217",
		      "value": 15229
		    }
		  ],
		  "free_asset_net_usageV2": [
		    {
		      "key": "1000217",
		      "value": 0
		    }
		  ]
		}
		```
		
### Account Bandwidth Information
- Returns bandwith information for the account. If a field doesn't appear, then the corresponding value is 0.
	- **POST** /v1/tron/wallet/getaccountnet
		-  Request

		``` json
		{
		  "address": "4170622A5F83CF2A5DE36433AD4F6CDEF078B01F05"
		}
		```
		
		-	Response

		``` json
		{
		  "freeNetLimit": 5000,
		  "assetNetUsed": [
		    {
		      "key": "1000217",
		      "value": 0
		    }
		  ],
		  "assetNetLimit": [
		    {
		      "key": "1000217",
		      "value": 0
		    }
		  ],
		  "TotalNetLimit": 43200000000,
		  "TotalNetWeight": 59143027
		}
		```

### Block by Block Height
- Returns the Block Object corresponding to the 'Block Height' specified (number of blocks preceding it).
	-	**POST** /v1/tron/wallet/getblockbynum	
		- Request

		``` json
		{
		  "num": 200
		}
		```
		
		-	Response

		``` json
		{
		  "blockID": "00000000000000c83ad3d3759a57db65c945053a0351bd044e3cb46dd939ecfd",
		  "block_header": {
		    "raw_data": {
		      "number": 200,
		      "txTrieRoot": "0000000000000000000000000000000000000000000000000000000000000000",
		      "witness_address": "41928c9af0651632157ef27a2cf17ca72c575a4d21",
		      "parentHash": "00000000000000c7d54cdd476da9f19493bfffe62757c9deee1b87342b898ea6",
		      "version": 6,
		      "timestamp": 1545846339000
		    },
		    "witness_signature": "68282eb82010fd38d4582da747764cd110a32c0a083bcadf0c0968ceb59a6711202ae3187d47c94f8535faba1f2e232c1b455f12b4101b91dacb40410806fef600"
		  }
		}
		```
		
### Block by ID
- Query block by ID.
	-	**POST** /v1/tron/wallet/getblockbyid
		- Request

		``` json
		{
		  "value": "00000000000000c83ad3d3759a57db65c945053a0351bd044e3cb46dd939ecfd"
		}
		```
		
		-	Response

		``` json
		{
		  "blockID": "00000000000000c83ad3d3759a57db65c945053a0351bd044e3cb46dd939ecfd",
		  "block_header": {
		    "raw_data": {
		      "number": 200,
		      "txTrieRoot": "0000000000000000000000000000000000000000000000000000000000000000",
		      "witness_address": "41928c9af0651632157ef27a2cf17ca72c575a4d21",
		      "parentHash": "00000000000000c7d54cdd476da9f19493bfffe62757c9deee1b87342b898ea6",
		      "version": 6,
		      "timestamp": 1545846339000
		    },
		    "witness_signature": "68282eb82010fd38d4582da747764cd110a32c0a083bcadf0c0968ceb59a6711202ae3187d47c94f8535faba1f2e232c1b455f12b4101b91dacb40410806fef600"
		  }
		}
		```
		
### Block by Latest Number
- Returns a list of block objects.
	-	**POST** /v1/tron/wallet/getblockbylatestnum	
		- Request

		``` json
		{
		  "num": 1
		}
		```
		
		-	Response

		``` json
		{
		  "block": [
		    {
		      "transactions": [
		        {
		          "raw_data": {
		            "ref_block_bytes": "d84e",
		            "ref_block_hash": "7fd027891175b96d",
		            "expiration": 1552896195000,
		            "contract": [
		              {
		                "type": "TriggerSmartContract",
		                "parameter": {
		                  "type_url": "type.googleapis.com/protocol.TriggerSmartContract",
		                  "value": "0a1541b9c2acd1dec60585e7894d63583dc41eef0c200312154134d03d7bd0b9bbd9bc742907bb52330ac1537acb22c402222bb7700000000000000000000000008ee29b4cf98777231756f3f50a545435a5055d1a00000000000000000000000000000000000000000000000000000000000000310000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000582f477a564c4b48692b4a47526e584d35473764514c4d426b5332655731777059644b4c79424a464f6e3352767167582b615634356c3634715367494270456448636c686246707641507648596b6668715672494461513d3d0000000000000000000000000000000000000000000000000000000000000000000000000000001c774b6e7870756179476d3354725a713679366b72367357736b53493d00000000"
		                }
		              }
		            ],
		            "timestamp": 1552896137306,
		            "fee_limit": 10000000
		          },
		          "signature": [
		            "51f7cd434c01ec6f87ca84d5eafccda815222d55ec5cb006e4f697c91fddc4e29c75d3ce95ca4fc5a6b41dd19507628e6272b6567365ab0bb488f20b4658e05700"
		          ],
		          "ret": [
		            {
		              "contractRet": "SUCCESS"
		            }
		          ]
		        },
		        {
		          "raw_data": {
		            "ref_block_bytes": "d84e",
		            "ref_block_hash": "7fd027891175b96d",
		            "expiration": 1552896195000,
		            "contract": [
		              {
		                "type": "TriggerSmartContract",
		                "parameter": {
		                  "type_url": "type.googleapis.com/protocol.TriggerSmartContract",
		                  "value": "0a154128d7f83009d3c6e26452b4cb1eef8dd931bc45d3121541e703ff48dbc60e41ef456ce5bd14d006213cb41d222420a98f6a000000000000000000000000000000000000000000000000000000005c8f508a"
		                }
		              }
		            ],
		            "timestamp": 1552896137939,
		            "fee_limit": 100000000
		          },
		          "signature": [
		            "0c36fee0583e6e39e3952f822e7631d60d8b3fa756310873d03ea89fbe7036c394e68cee974f63cae3225cc5891fff08c1e2ffa562320b22c02bb5bab09a09d000"
		          ],
		          "ret": [
		            {
		              "contractRet": "SUCCESS"
		            }
		          ]
		        },
		        {
		          "raw_data": {
		            "ref_block_bytes": "d84e",
		            "ref_block_hash": "7fd027891175b96d",
		            "expiration": 1552896195000,
		            "contract": [
		              {
		                "type": "TriggerSmartContract",
		                "parameter": {
		                  "type_url": "type.googleapis.com/protocol.TriggerSmartContract",
		                  "value": "0a1541d4025baca63d92e7bcf0b6ba34a14a65e28dc98b1215410d89b5bf6ea312dbc343e528900683b36975ed611880c2d72f2224359382580000000000000000000000000000000000000000000000000000000000000032"
		                }
		              }
		            ],
		            "timestamp": 1552896137933,
		            "fee_limit": 1000000000
		          },
		          "signature": [
		            "6ade644568d836089558b231827594ceb99e822bef9be27bfd9c790a442c2246c6b33cda7df991023066ab030fe02f2c620ca2bcc7f5d018dca02f2d7174770301"
		          ],
		          "ret": [
		            {
		              "contractRet": "SUCCESS"
		            }
		          ]
		        }
		      ],
		      "block_header": {
		        "raw_data": {
		          "timestamp": 1552896141000,
		          "txTrieRoot": "bcf977834904cff2249c8584e8e1f2ca1bd92f9d24e6e6782cab162ba8ed3bc2",
		          "parentHash": "000000000023d84f851ffb9d3139bd7cc259bf3b5a5953ace06f3ec46110d947",
		          "number": 2349136,
		          "witness_address": "41928c9af0651632157ef27a2cf17ca72c575a4d21",
		          "version": 7
		        },
		        "witness_signature": "65c253b94972c4b3c4da684fd28df3f85e0d21f2695e7250906a15fb1dce72fb2f51a99e30150f4e948b66eb84097c01a5f35048893a7b56e00146a04248271700"
		      }
		    }
		  ]
		}
		```
		
### Block by Limit Next
- Returns the list of Block Objects included in the 'Block Height' range specified.
	-	**POST** /v1/tron/wallet/getblockbylatestnum
		-	Request

		``` json
		{
		  "startNum": 1,
		  "endNum": 2
		}
		```
		
		-	Response
		
		``` json
		{
		  "block": [
		    {
		      "block_header": {
		        "raw_data": {
		          "timestamp": 1545845736000,
		          "txTrieRoot": "0000000000000000000000000000000000000000000000000000000000000000",
		          "parentHash": "000000000000000057945fc39376c3665d26a8743fc4c5c14f0fc7d8b00da8dd",
		          "number": 1,
		          "witness_address": "41928c9af0651632157ef27a2cf17ca72c575a4d21",
		          "version": 6
		        },
		        "witness_signature": "d69e24e843c03f0e549a33aeef65ecd55a3ac0aa9c37ec480a7a92e1e29bdcd025fbc4564cf5f95feae0aff3043c2e636eae5d72d66c0c12992ef55dc4c4a00901"
		      }
		    }
		  ]
		}
		```
		
### Get Latest Block
-	Query the latest block synced to the Full Node.
	-	**POST** /v1/tron/wallet/getnowblock
		-	Request

		``` json
		{
		  "req": "nowblock"
		}
		```
		
		-	Response
		
		``` json
		{
		  "blockID": "000000000023daf6aee54b4d88041fd9542d63119aaab95f79884fccd33c1985",
		  "block_header": {
		    "raw_data": {
		      "number": 2349814,
		      "txTrieRoot": "ce64fdf1e301de76c6b6d02ef3c4083bd20d42c02a687c9965994bb5fbaca2f2",
		      "witness_address": "41928c9af0651632157ef27a2cf17ca72c575a4d21",
		      "parentHash": "000000000023daf5158dea1b61cd1cec4551c9356887fdfd1b45cb0422e5de88",
		      "version": 7,
		      "timestamp": 1552898175000
		    },
		    "witness_signature": "718b6fbc6dbee6a6021dedfc37562c4030c1a7f75db71fe6524088243d500238088a90dfeac472a2e884870457c9fcd762525b36e82d26e19bc1283e0597057700"
		  },
		  "transactions": [
		    {
		      "ret": [
		        {
		          "contractRet": "SUCCESS"
		        }
		      ],
		      "signature": [
		        "c4d6c0fc5335da1a7665db99d794ec1938266e7f0cd8539ac8e0f9af49fbcd95d86c1e691c5f65b6a3449d6af4b397a21e60bebc2c2f271f7738996c72425cb300"
		      ],
		      "txID": "a37492f18bacd8a4d58106efa3ddef1e1f97db172f86b4b3cd541b3e839022e7",
		      "raw_data": {
		        "contract": [
		          {
		            "parameter": {
		              "value": {
		                "data": "20a98f6a000000000000000000000000000000000000000000000000000000005c8f587d",
		                "owner_address": "4128d7f83009d3c6e26452b4cb1eef8dd931bc45d3",
		                "contract_address": "41e703ff48dbc60e41ef456ce5bd14d006213cb41d"
		              },
		              "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
		            },
		            "type": "TriggerSmartContract"
		          }
		        ],
		        "ref_block_bytes": "daf5",
		        "ref_block_hash": "158dea1b61cd1cec",
		        "expiration": 1552898232000,
		        "fee_limit": 100000000,
		        "timestamp": 1552898173012
		      },
		      "raw_data_hex": "0a02daf52208158dea1b61cd1cec40c09dcaff982d5a8e01081f1289010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412540a154128d7f83009d3c6e26452b4cb1eef8dd931bc45d3121541e703ff48dbc60e41ef456ce5bd14d006213cb41d222420a98f6a000000000000000000000000000000000000000000000000000000005c8f587d70d4d0c6ff982d900180c2d72f"
		    },
		    {
		      "ret": [
		        {
		          "contractRet": "SUCCESS"
		        }
		      ],
		      "signature": [
		        "22beafa5494def53e1bc5d295528021bdfc6aa6f253c2400fb04cc820e05c7481c68b14b9482fe12918d40d86b2d1a4fdf3cb32b47ae475f1a9e0c532c6a502900"
		      ],
		      "txID": "4e4ef3c3c9ca240778cabd917bbbad75e6c8c59e1375b817a92c323229bdd8c4",
		      "raw_data": {
		        "contract": [
		          {
		            "parameter": {
		              "value": {
		                "data": "359382580000000000000000000000000000000000000000000000000000000000000032",
		                "owner_address": "41d4025baca63d92e7bcf0b6ba34a14a65e28dc98b",
		                "contract_address": "410d89b5bf6ea312dbc343e528900683b36975ed61",
		                "call_value": 100000000
		              },
		              "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
		            },
		            "type": "TriggerSmartContract"
		          }
		        ],
		        "ref_block_bytes": "daf5",
		        "ref_block_hash": "158dea1b61cd1cec",
		        "expiration": 1552898232000,
		        "fee_limit": 1000000000,
		        "timestamp": 1552898173386
		      },
		      "raw_data_hex": "0a02daf52208158dea1b61cd1cec40c09dcaff982d5a9301081f128e010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412590a1541d4025baca63d92e7bcf0b6ba34a14a65e28dc98b1215410d89b5bf6ea312dbc343e528900683b36975ed611880c2d72f222435938258000000000000000000000000000000000000000000000000000000000000003270cad3c6ff982d90018094ebdc03"
		    },
		    {
		      "ret": [
		        {
		          "contractRet": "SUCCESS"
		        }
		      ],
		      "signature": [
		        "6be055b73af61585a2ceccf7a07b59b7a69078dd7aadbf75d262e10195bbd4a87f87753bedad360a0f8422cb2d84870a72c1f14e18a7b7cc42a0e167ad50a86c00"
		      ],
		      "txID": "5cb538863680450508fdf6b8c0fde5ed99ddb43a1349b16b7c8fd8eb6c551aa0",
		      "raw_data": {
		        "contract": [
		          {
		            "parameter": {
		              "value": {
		                "data": "bb1086c10000000000000000000000000000000000000000000000000000000000000000",
		                "owner_address": "41d19cd0057df24d00df8f4bd23fcdda1b9ad018dd",
		                "contract_address": "41fd896be2b82a224118266ad6f073ffdf13dba238"
		              },
		              "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
		            },
		            "type": "TriggerSmartContract"
		          }
		        ],
		        "ref_block_bytes": "daf5",
		        "ref_block_hash": "158dea1b61cd1cec",
		        "expiration": 1552898232000,
		        "fee_limit": 1000000000,
		        "timestamp": 1552898173289
		      },
		      "raw_data_hex": "0a02daf52208158dea1b61cd1cec40c09dcaff982d5a8e01081f1289010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412540a1541d19cd0057df24d00df8f4bd23fcdda1b9ad018dd121541fd896be2b82a224118266ad6f073ffdf13dba2382224bb1086c1000000000000000000000000000000000000000000000000000000000000000070e9d2c6ff982d90018094ebdc03"
		    }
		  ]
		}
		```
		
### Broadcast Transaction
- Returns broadcast success or failure status.
	-	**POST** /v1/tron/wallet/broadcasttransaction	
		- Request

		``` json
		{
		  "raw_data": {
		    "ref_block_bytes": "e6dc",
		    "ref_block_hash": "cb95b9e4914854cc",
		    "expiration": 1552908215000,
		    "timestamp": 1552907315000,
		    "contract": [
		      {
		        "type": "TransferContract",
		        "parameter": {
		          "value": {
		            "amount": 50000000,
		            "owner_address": "412a00fc3f6068ed005d3d00b30609ddb79c26e213",
		            "to_address": "412be614e143e40979efd18b47cfe089c2ba5906f1"
		          },
		          "type_url": "type.googleapis.com/protocol.TransferContract"
		        }
		      }
		    ]
		  },
		  "signature": [
		    "edaa7bd24c09431ff423194f60730e421e5e207f747a8c1c738e370d2b779bf12677a3dc34b98afb6bc997d0b53175a4ac945dd7a571ec07138d243532840d4200"
		  ],
		  "txID": "1d3e4cfa27460a0b0ba3e3a8eaae61b8786979fef2cdbeea0df550fb60369e9d"
		}
		```
		
		-	Response
		
		``` json
		{
		  "result": true
		}
		```
		
### Validate Address
- Validates address, returns either true or false.
	-	**POST** /v1/tron/wallet/validateaddress
		- Request

		``` json
		{
		  "address": "4170622A5F83CF2A5DE36433AD4F6CDEF078B01F05"
		}
		```
		
		-	Response
		
		``` json
		{
		  "result": true,
		  "message": "Hex string format"
		}
		```
		
### Transactions by ID
- Query transaction by ID.
	-	**POST** /v1/tron/wallet/gettransactionbyid
		- Request

		```json
		{
		  "value": "d0807adb3c5412aa150787b944c96ee898c997debdc27e2f6a643c771edb5933"
		}
		```
		
		-	Response
		
		``` json
		{
		  "id": "9af66ca310ffb0f450c36128be690a65dc947eaeef424bb1512ad7b4e5958044",
		  "fee": 710890,
		  "blockNumber": 2349977,
		  "blockTimeStamp": 1552898664000,
		  "contractResult": [
		    "0000000000000000000000000000000000000000000000000000000005defda0000000000000000000000000000000000000000000000000000000000000001a"
		  ],
		  "contract_address": "414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		  "receipt": {
		    "energy_fee": 707700,
		    "energy_usage_total": 70770,
		    "net_fee": 3190,
		    "result": "SUCCESS"
		  },
		  "log": [
		    {
		      "address": "4d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		      "topics": [
		        "38498a37f911dcd45d3d0295f0ea4b117b4b2a1e3ecdc63bbc375e89f5366db7"
		      ],
		      "data": "0000000000000000000000000000000000000000000000000000000647aed498"
		    },
		    {
		      "address": "da84ca64e26fbaa0f756b23d36f0d1644e021708",
		      "topics": [
		        "0810ac8bfe739ca026451fbb5bdee58a52bdf5f8fb9aa321f52154e1481d06c7"
		      ],
		      "data": "000000000000000000000000a12ae80eeb9d006a64e95dd13df1dff76f7f4fc900000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000738e2eac"
		    },
		    {
		      "address": "4d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		      "topics": [
		        "e12d3a32e6efaf849a91ed7bbc5989d84662b78f949219c4b69cf35175dbb8b1",
		        "0000000000000000000000000000000000000000000000000000000000000002"
		      ],
		      "data": "000000000000000000000000a12ae80eeb9d006a64e95dd13df1dff76f7f4fc9000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000761d8694600000000000000000000000000000000000000000000000000000000005defda00000000000000000000000000000000000000000000000000000000002faf080"
		    }
		  ],
		  "internal_transactions": [
		    {
		      "hash": "da0588c6f39cbed2f65d810203b360f9ee1c6579b48b254ddaf78d0a719d7995",
		      "caller_address": "414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		      "transferTo_address": "41a12ae80eeb9d006a64e95dd13df1dff76f7f4fc9",
		      "callValueInfo": [
		        {
		          "callValue": 98500000
		        }
		      ],
		      "note": "63616c6c"
		    },
		    {
		      "hash": "be9c846156578536e7edede85810f072baa9e9517db04ec6314cb09f1b5cbac7",
		      "caller_address": "414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		      "transferTo_address": "41b3a1c3eb13837618249c16a5f300210aec7c4fa7",
		      "callValueInfo": [
		        {}
		      ],
		      "note": "63616c6c"
		    },
		    {
		      "hash": "40efc28bf57c6df16ebacd00194b2bee27680158cf712e438447db7ae50e3f42",
		      "caller_address": "414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		      "transferTo_address": "41de15a33d662153c69e4a2844f91437834f1f0496",
		      "callValueInfo": [
		        {}
		      ],
		      "note": "63616c6c"
		    },
		    {
		      "hash": "741e414373ba9c108485e4d288a1d98461bc2879adafa66d5f59ae68373696cc",
		      "caller_address": "41de15a33d662153c69e4a2844f91437834f1f0496",
		      "transferTo_address": "41da84ca64e26fbaa0f756b23d36f0d1644e021708",
		      "callValueInfo": [
		        {}
		      ],
		      "note": "63616c6c"
		    },
		    {
		      "hash": "8127b690804be49f46f678154157182a48008cd230cf8ead731240973ef28336",
		      "caller_address": "41de15a33d662153c69e4a2844f91437834f1f0496",
		      "transferTo_address": "41a09d5fd5fb19cfcc6f24972c4aa0f46da3c5f6d6",
		      "callValueInfo": [
		        {}
		      ],
		      "note": "63616c6c"
		    }
		  ]
		}
		```
		
### Transaction Info by ID
- Query transaction info by ID.
	-	**POST** /v1/tron/wallet/gettransactioninfobyid	
		- Request

		``` json
		{
		  "value": "d0807adb3c5412aa150787b944c96ee898c997debdc27e2f6a643c771edb5933"
		}
		```
		
		-	Response
		
		``` json
		{
		  "ret": [
		    {
		      "contractRet": "SUCCESS"
		    }
		  ],
		  "signature": [
		    "d273cb2238702927a41e2858b3ff31bf598ac307968dc211e9928bd91f53acf9c76eec211bea0c084dfdf17bac6678d6a1e3da043599de7c746e0d93f5d89f4f00"
		  ],
		  "txID": "9af66ca310ffb0f450c36128be690a65dc947eaeef424bb1512ad7b4e5958044",
		  "raw_data": {
		    "contract": [
		      {
		        "parameter": {
		          "value": {
		            "data": "53a04b050000000000000000000000000000000000000000000000000000000000000032",
		            "owner_address": "41a12ae80eeb9d006a64e95dd13df1dff76f7f4fc9",
		            "contract_address": "414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae",
		            "call_value": 50000000
		          },
		          "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
		        },
		        "type": "TriggerSmartContract"
		      }
		    ],
		    "ref_block_bytes": "db98",
		    "ref_block_hash": "21b5048fa06f8b60",
		    "expiration": 1552898721000,
		    "fee_limit": 1000000000,
		    "timestamp": 1552898661862
		  },
		  "raw_data_hex": "0a02db98220821b5048fa06f8b6040e889e8ff982d5a9301081f128e010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412590a1541a12ae80eeb9d006a64e95dd13df1dff76f7f4fc91215414d45db4d4e8e04f61cc73c2d3f41a1a7c9f50eae1880e1eb17222453a04b05000000000000000000000000000000000000000000000000000000000000003270e6bbe4ff982d90018094ebdc03"
		}
		```
		

### Get Contract
-	**POST** /v1/tron/wallet/getcontract
	-	Queries a contract's information from the blockchain. Returns SmartContract object.
		- Request

		``` json
		{
		  "value": "41909C9BFAB86CBF89B1A4D1488744111637719B14"
		}
		```
		
### Trigger Smart Contract
-	Returns TransactionExtention, which contains the unsigned Transaction.
	-	**POST** /v1/tron/wallet/triggersmartcontract
		- Request

		``` json
		{
		  "contract_address": "4189139CB1387AF85E3D24E212A008AC974967E561",
		  "function_selector": "set(uint256,uint256)",
		  "parameter": "00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
		  "fee_limit": 1000000,
		  "call_value": 100,
		  "owner_address": "41D1E7A6BC354106CB410E65FF8B181C600FF14292"
		}
		```
		
		-	Response

		``` json
		{
		  "result": {
		    "result": true
		  },
		  "constant_result": [
		    "0000000000000000000000000000000000000000000000000000000041e4ddb6"
		  ],
		  "transaction": {
		    "ret": [
		      {}
		    ],
		    "txID": "454246733416b780411e5c59618264e3a881e5eee5858894f3a625a8342a6ea7",
		    "raw_data": {
		      "contract": [
		        {
		          "parameter": {
		            "value": {
		              "data": "70a08231000000000000000000000000c12fb0805f3ba3ca3a01d258a887dd3e4db731a6",
		              "owner_address": "41c12fb0805f3ba3ca3a01d258a887dd3e4db731a6",
		              "contract_address": "414b9223269e1c53b6b72745e3332623c21e70b805"
		            },
		            "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
		          },
		          "type": "TriggerSmartContract"
		        }
		      ],
		      "ref_block_bytes": "e765",
		      "ref_block_hash": "62c12bc2214f8603",
		      "expiration": 1552907784000,
		      "fee_limit": 1000000,
		      "timestamp": 1552907726027
		    },
		    "raw_data_hex": "0a02e765220862c12bc2214f860340c09e9184992d5a8e01081f1289010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412540a1541c12fb0805f3ba3ca3a01d258a887dd3e4db731a61215414b9223269e1c53b6b72745e3332623c21e70b805222470a08231000000000000000000000000c12fb0805f3ba3ca3a01d258a887dd3e4db731a670cbd98d84992d9001c0843d"
		  }
		}
		```
---------------------------------------
# CYBAVO Litecoin API Sample

### LTC - Transaction Fee 
- Query transaction fee
	- **GET** /v1/ltc/wallet/fee
		- Response

		``` json
		{
			"error_code": 0,
  		"result": {
    		"High": "0.00257332",
    		"Medium": "0.00257332",
    		"Low": "0.00257332",
    		"Suggest": "0.00257332"
  		}
		}
		```

### LTC - Get address balance 
- Query balance in specific address
	- **GET** /v1/ltc/wallet/addressbalance/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Balance": "4.99865"
  		}
		}
		```

### LTC - Get confirmation of TXID 
- Query confirmation number in specific txid
	- **GET**  /v1/ltc/transaction/confirm/`LTC_TXID`
		- Response
		
		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Number": 3580
  		}
		}
		```

### LTC - Get unconfirmed balance 
- Query unconfirmed balance in specific address
	- **GET**  /v1/ltc/transaction/addressunconfirmbalance/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"UnconfirmBalance": "0.1234"
  		}
		}
		```

### LTC - Create raw transaction 
- Create transaction information. Server selects eligible UTXO and returns encoding string of transaction
	- **POST**  /v1/ltc/transaction/createpayment
		- Request

		``` json
		{
  		"FromAddress": "QNdiW2GFR4mTw1in7Nn2LAEV1NeFphUAJg",
  		"ToAddress": "miQihDe6EvUzaKX7DB4yFfTEB3P6FYDVWK",
  		"Amount": "0.87788769",
  		"Fee": "0.0000416"
		}
		```

		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"TransactionHash": "7b227261777478223a22303130303030303030313863613563373065356530383331396538303161666434313466616237366362613131323237643732653833316535653461613730633335633766313762633030313030303030303030666666666666666630326531386333623035303030303030303031393736613931343166626138373432653562343764363363313030323732653364303137633739373264643339323638386163623435376332303030303030303030303137613931343136343430646564663030323066386561366564353234313636633530646238353835633233333838373030303030303030222c22696e707574616d6f756e74223a5b7b2261646472657373223a22514e64695732474652346d547731696e374e6e324c414556314e6546706855414a67222c22616d6f756e74223a3130303533303230357d5d7d",
    		"Fee": "5e-05"
  		}
		}
		```

### LTC - Send signed transaction 
- Send signed transaction to blockchain
	- **POST**  /v1/ltc/transaction/submitpayment
		- Request
		
		``` json
		{
  		"Transaction": "01000000000101fbb28dd706c6817525a55835cbd8da78b67f617f731dbd0fbbd98460492d543d01000000171600148dc0222aad18401c9b5a510b10a3177729e6ccc0ffffffff0210270000000000001976a9141fba8742e5b47d63c100272e3d017c7972dd392688ac101bcb1d0000000017a914aed75052d9ecaf0a2562aaee8ee044dfc2bb75bd8702473044022035a599e87c928a6e27310e63c2e9e6c7b086504afb7ee44c9baa442d123375e20220023fd88234fce894484a88061cb8b213c272326c41f8d552d10055043d885476012103379b73388a9cc5a09f310d96cc357a50d1ebbe57e0ef904cf205714411eacac100000000"
		}
		```

		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"State": true,
    		"TxId": "ea6bff6609824a9970369dbc9360dd0e9bc1efba4794792fe20c4c6e14a72b37"
  		}
		}
		```

### LTC - Get UTXO in address 
- Query UTXO in specific address
	- **GET**  /v1/ltc/transaction/addressutxo/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": [
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "5feb587d6e1df471bbda672cad8a7431f1e5c99a4b5299e9d3f3561296c03e08",
      		"outputIndex": 0,
      		"satoshis": 998009364,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "5feb587d6e1df471bbda672cad8a7431f1e5c99a4b5299e9d3f3561296c03e08",
      		"outputIndex": 1,
      		"satoshis": 1000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "06d1d92052df90081299bf0275c9d3c82b518ee699723b1061f302a10a8609f2",
      		"outputIndex": 0,
      		"satoshis": 50000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "025fc20f8fc9794c117b24943ca8458cdd804bdf4713d93ac1b72a75818f8110",
      		"outputIndex": 0,
      		"satoshis": 500000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1025276,
      		"confirmations": 821,
      		"pending": false
    		}
  		]
		}
		```

### LTC - Get transaction history in address 
- Query transaction history in address by specific number
	- **GET**  /v1/ltc/wallet/addresshistory/`LTC_ADDRESS`/`NUMBER`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"txids": [
      		"3d542d496084d9bb0fbd1d737f617fb678dad8cb3558a5257581c606d78db2fb",
      		"d032a022140032c3b621ad4adead93a50b764fd457b0ff2a5d12799fec54304b",
      		"6fd5ee0b912765a2d911f7fad0fcf189d3c982d85730d127f6acc9678856390e",
      		"d81f95db9d107121f5ca799a42c1dd7c706aefda461257f9f766bc9c1d0a2dc0",
      		"632bfc5fc645f3cdea27a6f2ddcaa758d18ef918fb06c6725d1980efc6f7f396",
      		"62153b0636f88e3149b1bd5ecd0bf5bb69fb50b44cd8a52160d182115efd8b07",
      		"64ffbe475f74cce4a71149914316ffc9f5503fc7d2ee414cba85672e4e12d9f5",
      		"0b3b76b03b4cff268947af3ec1dea4cbdf52794c5fbe4a7934ac3d1efbe938b0",
      		"d5604f5c644700ca6955385f5e558c00175559b036237479bd93e8bfef5c5801",
      		"844ecfff4a218f0c09c60f631d64ca0fbb57ea17220adf4a6e16391a3e10f3e0"
    		],
    		"txid_states": null,
    		"txcount": 10,
    		"txin_count": 0,
    		"txout_count": 0,
    		"address": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT"
  		}
		}
		```

### LTC - Get transaction information 
- Query transaction information in specific TXID and address
	- **GET**  /v1/ltc/transaction/txinfo/`LTC_TXID`/`ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    	"txid": "fe2ff4efe9aad684059ea135b7630be571ada613e1ef88a995253f29592fceee",
    	"from": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT",
    	"to": "miQihDe6EvUzaKX7DB4yFfTEB3P6FYDVWK",
    	"amount": "0.0001",
    	"fee": "0.00005",
    	"timestamp": 1553498925,
    	"out": true,
    	"token": "",
    	"status": 1
  		}
		}
		```

### LTC - Get batch of address balance 
- Query batch of address balance. The maximum number of address is 20
	- **POST**  /v1/ltc/wallet/addressbalance
		- Request
		``` json
		{
  		"Addresses": [
    		{
      		"Address": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT"
    		},
    		{
      		"Address": "QaAqKiTwm5qpYyjuSLRXhuAHtpBuWn6vFU"
    		}
  		]
		}
		```
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Balances": [
      		{
        		"Address": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT",
        		"Token": "",
        		"Balance": "4.99715"
      		},
      		{
        		"Address": "QaAqKiTwm5qpYyjuSLRXhuAHtpBuWn6vFU",
        		"Token": "",
        		"Balance": "159472.49358268"
      		}
    		]
  		}
		}
		```

### LTC - Get balance list of XPub
- Query balance list of XPub
	- **GET**  /v1/ltc/wallet/xpubbalance/`LTC_XPUB`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Balances": [
      		{
        		"Address": "n15ebjhpvs9Xr2H7R6WDdfBeAx7kEdi9CH",
        		"Token": "",
        		"Balance": "1.0004"
      		},
      		{
        		"Address": "mkJ6dFu4pzohktR8rwH8q1eK3qs8BTsBsY",
        		"Token": "",
        		"Balance": "4.9949"
      		}
    		]
  		}
		}
		```

### LTC - Create transaction blob with XPub
- Create transaction blob with XPub. UTXOs in this blob include all addresses which created by this XPub. Request must specify next address (ReceivingAddress) of XPub.
	- **POST**  /v1/ltc/transaction/xpubcreatepayment
		- Request

		``` json
		{
  		"Type": 0,
  		"ToAddress": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT",
  		"Amount": "0.0001",
  		"Fee": "0.0000416",
  		"XPub": "tpubDH9DBHzyDeLhoXWfyaaWFjiDRXsYguxZHKmbqTXqRMsHyt21jA1cXhMCjuBvFo2Q6PzY4UwRS1zArQH1aaKADdASW3vyZWwvRSF9vVLvmrL",
  		"ReceivingAddress": "mkJ6dFu4pzohktR8rwH8q1eK3qs8BTsBsY"
		}
		```

		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"TransactionHash": "7b227261777478223a22303130303030303030316465333064343633633562353132363830643835613366323331306563623363326263313865623264346664623664636334643863346430366562303936363230313030303030303030666666666666666630323130323730303030303030303030303031376139313461656437353035326439656361663061323536326161656538656530343464666332626237356264383733383632633531643030303030303030313937366139313433343661326430353836373736383730336531363633643161633364303833366634343730666338383861633030303030303030222c22696e707574616d6f756e74223a5b7b2261646472657373223a226d6b4a3664467534707a6f686b745238727748387131654b33717338425473427359222c22616d6f756e74223a3439393439303030307d5d7d",
    		"Fee": "5e-05"
  		}
		}
		```


---------------------------------------
# CYBAVO Bitcoin API Sample

### BTC - Transaction Fee 
- Query transaction fee
	- **GET** /v1/btc/wallet/fee
		- Response

		``` json
		{
			"error_code": 0,
  		"result": {
    		"High": "0.00257332",
    		"Medium": "0.00257332",
    		"Low": "0.00257332",
    		"Suggest": "0.00257332"
  		}
		}
		```

### BTC - Get address balance 
- Query balance in specific address
	- **GET** /v1/btc/wallet/addressbalance/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Balance": "4.99865"
  		}
		}
		```

### BTC - Get confirmation of TXID 
- Query confirmation number in specific txid
	- **GET**  /v1/btc/transaction/confirm/`LTC_TXID`
		- Response
		
		``` json
		{
  		"error_code": 0,
  		"result": {
    		"Number": 3580
  		}
		}
		```

### BTC - Get unconfirmed balance 
- Query unconfirmed balance in specific address
	- **GET**  /v1/btc/transaction/addressunconfirmbalance/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"UnconfirmBalance": "0.1234"
  		}
		}
		```

### BTC - Create raw transaction 
- Create transaction information. Server selects eligible UTXO and returns encoding string of transaction
	- **POST**  /v1/btc/transaction/createpayment
		- Request

		``` json
		{
  		"FromAddress": "QNdiW2GFR4mTw1in7Nn2LAEV1NeFphUAJg",
  		"ToAddress": "miQihDe6EvUzaKX7DB4yFfTEB3P6FYDVWK",
  		"Amount": "0.87788769",
  		"Fee": "0.0000416"
		}
		```

		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"TransactionHash": "7b227261777478223a22303130303030303030313863613563373065356530383331396538303161666434313466616237366362613131323237643732653833316535653461613730633335633766313762633030313030303030303030666666666666666630326531386333623035303030303030303031393736613931343166626138373432653562343764363363313030323732653364303137633739373264643339323638386163623435376332303030303030303030303137613931343136343430646564663030323066386561366564353234313636633530646238353835633233333838373030303030303030222c22696e707574616d6f756e74223a5b7b2261646472657373223a22514e64695732474652346d547731696e374e6e324c414556314e6546706855414a67222c22616d6f756e74223a3130303533303230357d5d7d",
    		"Fee": "5e-05"
  		}
		}
		```

### BTC - Send signed transaction 
- Send signed transaction to blockchain
	- **POST**  /v1/btc/transaction/submitpayment
		- Request
		
		``` json
		{
  		"Transaction": "01000000000101fbb28dd706c6817525a55835cbd8da78b67f617f731dbd0fbbd98460492d543d01000000171600148dc0222aad18401c9b5a510b10a3177729e6ccc0ffffffff0210270000000000001976a9141fba8742e5b47d63c100272e3d017c7972dd392688ac101bcb1d0000000017a914aed75052d9ecaf0a2562aaee8ee044dfc2bb75bd8702473044022035a599e87c928a6e27310e63c2e9e6c7b086504afb7ee44c9baa442d123375e20220023fd88234fce894484a88061cb8b213c272326c41f8d552d10055043d885476012103379b73388a9cc5a09f310d96cc357a50d1ebbe57e0ef904cf205714411eacac100000000"
		}
		```

		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"State": true,
    		"TxId": "ea6bff6609824a9970369dbc9360dd0e9bc1efba4794792fe20c4c6e14a72b37"
  		}
		}
		```

### BTC - Get UTXO in address 
- Query UTXO in specific address
	- **GET**  /v1/btc/transaction/addressutxo/`LTC_ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": [
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "5feb587d6e1df471bbda672cad8a7431f1e5c99a4b5299e9d3f3561296c03e08",
      		"outputIndex": 0,
      		"satoshis": 998009364,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "5feb587d6e1df471bbda672cad8a7431f1e5c99a4b5299e9d3f3561296c03e08",
      		"outputIndex": 1,
      		"satoshis": 1000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "06d1d92052df90081299bf0275c9d3c82b518ee699723b1061f302a10a8609f2",
      		"outputIndex": 0,
      		"satoshis": 50000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1024468,
      		"confirmations": 1629,
      		"pending": false
    		},
    		{
      		"address": "QTMdUhWACZxfxKTC9jFGLJ6QJCQKzz4jea",
      		"txid": "025fc20f8fc9794c117b24943ca8458cdd804bdf4713d93ac1b72a75818f8110",
      		"outputIndex": 0,
      		"satoshis": 500000000,
      		"script": "a9144a11d5c96c66a94e0bccc6f81b1244013bcf2b3687",
      		"height": 1025276,
      		"confirmations": 821,
      		"pending": false
    		}
  		]
		}
		```

### BTC - Get transaction history in address 
- Query transaction history in address by specific number
	- **GET**  /v1/btc/wallet/addresshistory/`LTC_ADDRESS`/`NUMBER`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"txids": [
      		"3d542d496084d9bb0fbd1d737f617fb678dad8cb3558a5257581c606d78db2fb",
      		"d032a022140032c3b621ad4adead93a50b764fd457b0ff2a5d12799fec54304b",
      		"6fd5ee0b912765a2d911f7fad0fcf189d3c982d85730d127f6acc9678856390e",
      		"d81f95db9d107121f5ca799a42c1dd7c706aefda461257f9f766bc9c1d0a2dc0",
      		"632bfc5fc645f3cdea27a6f2ddcaa758d18ef918fb06c6725d1980efc6f7f396",
      		"62153b0636f88e3149b1bd5ecd0bf5bb69fb50b44cd8a52160d182115efd8b07",
      		"64ffbe475f74cce4a71149914316ffc9f5503fc7d2ee414cba85672e4e12d9f5",
      		"0b3b76b03b4cff268947af3ec1dea4cbdf52794c5fbe4a7934ac3d1efbe938b0",
      		"d5604f5c644700ca6955385f5e558c00175559b036237479bd93e8bfef5c5801",
      		"844ecfff4a218f0c09c60f631d64ca0fbb57ea17220adf4a6e16391a3e10f3e0"
    		],
    		"txid_states": null,
    		"txcount": 10,
    		"txin_count": 0,
    		"txout_count": 0,
    		"address": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT"
  		}
		}
		```

### BTC - Get transaction information 
- Query transaction information in specific TXID and address
	- **GET**  /v1/btc/transaction/txinfo/`LTC_TXID`/`ADDRESS`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    	"txid": "fe2ff4efe9aad684059ea135b7630be571ada613e1ef88a995253f29592fceee",
    	"from": "QcYTcwUNkKu4e6BJWZTYeT6NR16Z9C72FT",
    	"to": "miQihDe6EvUzaKX7DB4yFfTEB3P6FYDVWK",
    	"amount": "0.0001",
    	"fee": "0.00005",
    	"timestamp": 1553498925,
    	"out": true,
    	"token": "",
    	"status": 1
  		}
		}
		```

### BTC - Get block information 
- Query block information in specific block number
	- **GET**  /v1/btc/block/blockinfo/`BTC_BLOCK_NUMBER`
		- Response

		``` json
		{
  		"error_code": 0,
  		"result": {
    		"currency_type": 0,
    		"btc_block": {
      		"hash": "000000000001ef600f2273987ffd44880da73a99af755f05cf3f145aad276734",
      		"confirmations": 917203,
      		"strippedsize": 533,
      		"size": 533,
      		"weight": 2132,
      		"height": 568718,
      		"version": 3,
      		"versionHex": "00000003",
      		"merkleroot": "d41ef19eadde14e961b61a68883f3d1fba60be2b6bae23cf0f202ce2c9f3b144",
      		"tx": [
        
      		],
      		"time": 1442578724,
      		"nonce": 2864936812,
      		"bits": "1b0a38f0",
      		"difficulty": 6410.913520097442,
      		"previousblockhash": "0000000000044e2c0ced6b945bb32c2833d39d6388ffef1f833e709718e29b77",
      		"nextblockhash": "0000000000005cf47cee047df43d4c4854cd1b9dddf80de96b9fd9766370a66b"
    		},
    		"eth_block": null,
    		"eth_receipt": null
  		}
		}
		```
