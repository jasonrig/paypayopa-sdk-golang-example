package main

import (
	"github.com/google/uuid"
	"github.com/jasonrig/paypayopa-sdk-golang/api"
	"github.com/jasonrig/paypayopa-sdk-golang/api/v2/codes"
	"github.com/jasonrig/paypayopa-sdk-golang/request"
	"log"
	"net/http"
)

func main() {
	// This will obtain credentials from your system environment
	// Be sure to set PAYPAY_CLIENT_ID and PAYPAY_CLIENT_SECRET
	auth, err := request.NewAuth(nil, nil)
	if err != nil {
		log.Panic(err)
	}

	// Define QR Code creation request parameters
	paymentId, _ := uuid.NewRandom()
	orderDescription := "This is your order description"
	requestPayload := codes.PostPayload{
		MerchantPaymentId: paymentId.String(),
		Amount: api.Amount{
			Amount:   100,
			Currency: "JPY",
		},
		CodeType:         "ORDER_QR",
		OrderDescription: &orderDescription,
		OrderItems: &[]api.OrderItems{
			{
				Name: "Fun thing",
				UnitPrice: &api.Amount{
					Amount:   50,
					Currency: "JPY",
				},
				Quantity: 1,
			},
			{
				Name: "Another fun thing",
				UnitPrice: &api.Amount{
					Amount:   50,
					Currency: "JPY",
				},
				Quantity: 1,
			},
		},
	}

	// Set up the request
	apiRequest := &codes.Post{
		// Choose an execution environment
		// - api.SandboxEnvironment
		// - api.StagingEnvironment
		// - api.ProductionEnvironment
		Environment: api.SandboxEnvironment,
		Payload:     requestPayload,
	}

	// Execute the API request and collect the response
	response := &codes.PostResponse{}
	err = apiRequest.MakeRequest().Call(auth, &http.Client{}, response)

	// Display the result
	if err != nil {
		log.Panicf("API call returned an error, %s", err)
	} else {
		log.Printf("QR Code URL: %s", *response.Url)
	}
}
