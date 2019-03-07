# CYBABO EOS API Sample

### Prerequisite

- Please obtain following info from CYBAVO Team
	- API URL 
	- API code
	- API secret

## API set

### Account
- /v1/eos/account/`EOS_ACCOUNT_NAME`/info
	- 查詢特定帳戶的基本訊息
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

- /v1/eos/account/`EOS_ACCOUNT_NAME`/resouce
	- 查詢特定帳戶的資源訊息
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

- /v1/eos/account/`EOS_ACCOUNT_NAME`/balance
	- 查詢特定帳戶的餘額訊息
		- Response

		``` json
		{
		  "error_code": 0,
		  "result": [
		    "2.3500 EOS"
		  ]
		}
		```

### Block
- /v1/eos/block/`BLOCK_NUM`
	- 查詢特定區塊內容包含區塊高度及區塊中交易的詳細訊息
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
	
- /v1/eos/block/latest
	- 獲取最新區塊訊息包含區塊高度及區塊中交易的詳細訊息
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

- /v1/eos/block/latest\_irreversible
	- 獲取最新不可逆區塊訊息包含區塊高度及區塊中交易的詳細訊息
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
		
- /v1/eos/proxy/chain/abi\_json\_to\_bin
	-  把json轉換為binary
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
		
- /v1/eos/proxy/chain/push\_transaction
	-  接收一簽名後的JSON格式的交易並廣播道區塊鏈上
		-  Request

		``` json
		{
		  "signatures": [
		    "SIG_K1_K4gsBzrZ5dTPrK2dv1bvwttcA7aTuFFyi4X43NDPPxExLvnDxGFpkHx8tmte22sEMKgopcBYT7dvoZgVJ7HFpyQJsrZDuo"
		  ],
		  "compression": "none",
		  "packed_context_free_data": "",
		  "packed_trx": "897ef15ba927136993dd000000000100a6823403ea3055000000572d3ccdcd0190dd39e69a64a64100000000a8ed32322190dd39e69a64a641e05b3597d15cfd45640000000000000004454f530000000000000000000000000000000000000000000000000000000000000000000000000000"
		}
		```