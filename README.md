## Receipt Processor

A receipt processor as described in [this](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main) repository

# To run server
1. `git clone https://github.com/jessyec-s/receipt-processor-challenge.git`
2. `cd receipt-processor-challenge`
3. `go run .`


# Supported Requests

`GET /receipts/{id}/points`

- Returns number of points for given receipt ID

`POST /receipts/process` with receipt payload

- Returns ID of receipt
