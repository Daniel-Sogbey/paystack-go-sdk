package Transactions

import (
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestInitialize(t *testing.T) {

	client := paystack.NewClient("sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb", "application/json")

	sampleInitializeTransactionRequest := &InitializeTransactionRequest{
		Email:  "1@2.com",
		Amount: "20",
	}

	response, err := Initialize(client, sampleInitializeTransactionRequest)

	if err != nil {
		t.Errorf("Expected response of status %v, but got a status of %v and an error that says %v", response.Status, response.Status, err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}

}