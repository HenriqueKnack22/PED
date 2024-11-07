package main

import (
	"net/rpc"
	"testing"
)

const (
	serverAddress = "localhost:5000" 
)

type ConsultBalanceRequest struct {
	Name string
}

type ConsultBalanceResponse struct {
	Balance float64
}

func connectRPC() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", serverAddress) 
	if err != nil {
		return nil, err
	}
	return client, nil
}

func setupClient(t *testing.T) *rpc.Client {
	client, err := connectRPC() 
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	return client
}

func CreateAccount(client *rpc.Client, name string, balance float64) error {
	return nil
}

func Deposit(client *rpc.Client, name string, amount float64) error {
	return nil
}

func Withdraw(client *rpc.Client, name string, amount float64) error {
	return nil
}

func DeleteAccount(client *rpc.Client, name string) error {
	return nil
}

func TestCreateAccount(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	name := "testUser"
	balance := 100.0
	err := CreateAccount(client, name, balance)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}

	var response ConsultBalanceResponse
	err = client.Call("Accounts.ConsultBalance", ConsultBalanceRequest{Name: name}, &response)
	if err != nil {
		t.Errorf("Failed to consult balance: %v", err)
	}
	if response.Balance != balance {
		t.Errorf("Expected balance %.2f, got %.2f", balance, response.Balance)
	}
}

func TestDeposit(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	name := "testUser"
	initialBalance := 100.0
	depositAmount := 50.0
	err := CreateAccount(client, name, initialBalance)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	err = Deposit(client, name, depositAmount)
	if err != nil {
		t.Fatalf("Failed to deposit amount: %v", err)
	}

	var response ConsultBalanceResponse
	err = client.Call("Accounts.ConsultBalance", ConsultBalanceRequest{Name: name}, &response)
	if err != nil {
		t.Errorf("Failed to consult balance after deposit: %v", err)
	}
	expectedBalance := initialBalance + depositAmount
	if response.Balance != expectedBalance {
		t.Errorf("Expected balance %.2f after deposit, got %.2f", expectedBalance, response.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	name := "testUserWithdraw"
	initialBalance := 200.0
	withdrawAmount := 50.0
	err := CreateAccount(client, name, initialBalance)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	err = Withdraw(client, name, withdrawAmount)
	if err != nil {
		t.Fatalf("Failed to withdraw amount: %v", err)
	}

	var response ConsultBalanceResponse
	err = client.Call("Accounts.ConsultBalance", ConsultBalanceRequest{Name: name}, &response)
	if err != nil {
		t.Errorf("Failed to consult balance after withdrawal: %v", err)
	}
	expectedBalance := initialBalance - withdrawAmount
	if response.Balance != expectedBalance {
		t.Errorf("Expected balance %.2f after withdrawal, got %.2f", expectedBalance, response.Balance)
	}
}

func TestDeleteAccount(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	name := "testUserDelete"
	err := CreateAccount(client, name, 50.0)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	err = DeleteAccount(client, name)
	if err != nil {
		t.Fatalf("Failed to delete account: %v", err)
	}

	var response ConsultBalanceResponse
	err = client.Call("Accounts.ConsultBalance", ConsultBalanceRequest{Name: name}, &response)
	if err == nil {
		t.Error("Expected error when consulting balance of deleted account, but got none.")
	}
}
