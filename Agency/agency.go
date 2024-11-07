//codigo que simula a agencia

package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

type CreateAccountRequest struct {
	Balance float64
	Name    string
	Id      string
}

type CreateAccountResponse struct {
	Message string
}

type ConsultBalanceRequest struct {
	Name string
}

type ConsultBalanceResponse struct {
	Balance float64
}

type DeleteAccountRequest struct {
	Name string
	Id   string
}

type DeleteAccountResponse struct {
	Message string
}

const agencyId = "7777"

func connectRPC(machine string, port int) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", machine, port))
	if err != nil {
		return nil, fmt.Errorf("connecting to server: %w", err)
	}
	return client, nil
}

func CreateAccount(client *rpc.Client, name string, balance float64) {
	request := CreateAccountRequest{Name: name, Balance: balance, Id: agencyId}
	var response CreateAccountResponse

	err := client.Call("Accounts.CreateAccount", request, &response)
	if err != nil {
		fmt.Println("Error: Creating account:", err)
	} else {
		fmt.Printf("Name: %s\nBalance: %.2f\nMessage: %s\n", name, balance, response.Message)
	}
}

func ConsultBalance(client *rpc.Client, name string) {
	request := ConsultBalanceRequest{Name: name}
	var response ConsultBalanceResponse

	err := client.Call("Accounts.ConsultBalance", request, &response)
	if err != nil {
		fmt.Println("Error: Consult balance:", err)
	} else {
		fmt.Printf("Name: %s\nBalance: %.2f\n", name, response.Balance)
	}
}

func DeleteAccount(client *rpc.Client, name string) {
	request := DeleteAccountRequest{Name: name, Id: agencyId}
	var response DeleteAccountResponse

	err := client.Call("Accounts.DeleteAccount", request, &response)
	if err != nil {
		fmt.Println("Error: Delete account:", err)
	} else {
		fmt.Printf("Name: %s\nMessage: %s\n", name, response.Message)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run agency.go <command> <machine> <port> <name> [balance]")
		return
	}

	command, machine := os.Args[1], os.Args[2]
	port, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error converting port:", err)
		return
	}

	name := os.Args[4]

	client, err := connectRPC(machine, port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer client.Close()

	switch command {
	case "create":
		if len(os.Args) < 6 {
			fmt.Println("Usage: go run agency.go create <machine> <port> <name> <balance>")
			return
		}
		balance, err := strconv.ParseFloat(os.Args[5], 64)
		if err != nil {
			fmt.Println("Error converting balance:", err)
			return
		}
		CreateAccount(client, name, balance)
	case "consult":
		ConsultBalance(client, name)
	case "delete":
		DeleteAccount(client, name)
	default:
		fmt.Println("Unknown command:", command)
	}
}
