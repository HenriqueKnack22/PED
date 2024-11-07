//codigo do caixa/teller.

package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

type DepositRequest struct {
	Name   string
	Amount float64
}

type DepositResponse struct {
	Message string
}

type WithdrawRequest struct {
	Name   string
	Amount float64
}

type WithdrawResponse struct {
	Message string
}

type ConsultBalanceRequest struct {
	Name string
}

type ConsultBalanceResponse struct {
	Balance float64
}

func parseAmount(amountString string) (float64, error) {
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		return 0, fmt.Errorf("converting amount: %w", err)
	}
	return amount, nil
}

func connectRPC(machine string, port int) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", machine, port))
	if err != nil {
		return nil, fmt.Errorf("connecting to server: %w", err)
	}
	return client, nil
}

func Deposit(client *rpc.Client, name string, amount float64) {
	request := DepositRequest{Name: name, Amount: amount}
	var response DepositResponse

	err := client.Call("Accounts.Deposit", request, &response)
	if err != nil {
		fmt.Println("Error depositing amount:", err)
	} else {
		fmt.Println(response.Message)
	}
}

func Withdraw(client *rpc.Client, name string, amount float64) {
	request := WithdrawRequest{Name: name, Amount: amount}
	var response WithdrawResponse

	err := client.Call("Accounts.Withdraw", request, &response)
	if err != nil {
		fmt.Println("Error withdrawing amount:", err)
	} else {
		fmt.Println(response.Message)
	}
}

func ConsultBalance(client *rpc.Client, name string) {
	request := ConsultBalanceRequest{Name: name}
	var response ConsultBalanceResponse

	err := client.Call("Accounts.ConsultBalance", request, &response)
	if err != nil {
		fmt.Println("Error consulting balance:", err)
	} else {
		fmt.Printf("Name: %s\nBalance: %.2f\n", name, response.Balance)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run teller.go <command> <machine> <port> <name> [amount]")
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
	case "deposit":
		if len(os.Args) < 6 {
			fmt.Println("Usage: go run teller.go deposit <machine> <port> <name> <amount>")
			return
		}
		amount, err := parseAmount(os.Args[5])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		Deposit(client, name, amount)
	case "withdraw":
		if len(os.Args) < 6 {
			fmt.Println("Usage: go run teller.go withdraw <machine> <port> <name> <amount>")
			return
		}
		amount, err := parseAmount(os.Args[5])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		Withdraw(client, name, amount)
	case "consult":
		ConsultBalance(client, name)
	default:
		fmt.Println("Unknown command:", command)
	}
}
