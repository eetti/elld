package blockchain

// GenesisBlock describes the first block
const GenesisBlock = `{
    "header": {
        "parentHash": null,
        "creatorPubKey": "49VzsGezoNjJQxHjekoCQP9CXZUs34CmCY53kGaHyR9rCJQJbJW",
        "number": 1,
        "stateRoot": "abc",
        "transactionRoot": "jsjhf9e3i9nfi",
        "nonce": 0,
        "mixHash": "",
        "difficulty": "500000",
        "timestamp": 0
    },
    "transactions": [
        {
            "type": 1,
            "nonce": 2,
            "to": "efjshfhh389djn29snmnvuis",
            "senderPubKey": "xsj2909jfhhjskmj99k",
            "value": "100.333",
            "timestamp": 1508673895,
            "fee": "0.00003",
            "invokeArgs": null,
            "sig": "93udndte7hxbvhivmnzbzguruhcbybcdbxcbyulmxsncs",
            "hash": "93udndte7hxbvhivmnzbzguruhcbybcdbxcbyulmxsncs"
        },
        {
            "type": 2,
            "nonce": 2,
            "to": null,
            "senderPubKey": "48d9u6L7tWpSVYmTE4zBDChMUasjP5pvoXE7kPw5HbJnXRnZBNC",
            "value": "10",
            "timestamp": 1508673895,
            "fee": "0.00003",
            "invokeArgs": null,
            "sig": "93udndte7hxbvhivmnzbzguruhcbybcdbxcbyulmxsncs",
            "hash": "shhfd7387ydhudhsy8ehhhfjsg748hd"
        }
    ],
    "hash": "lkssjajsdnaskomcskosks",
    "sig": "abc"
}`
