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

- /v1/eos/block/latest_irreversible
	- 獲取最新不可逆區塊訊息包含區塊高度及區塊中交易的詳細訊息
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
