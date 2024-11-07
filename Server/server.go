//Codigo que vai rodar o servidor.

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type Account struct {
	Name    string
	Balance float64
}

type Accounts struct {
	mu       sync.Mutex
	accounts map[string]*Account
}

type CreateAccountRequest struct {
	Name    string
	Balance float64
}

type CreateAccountResponse struct {
	Message string
}

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

type DeleteAccountRequest struct {
	Name string
}

type DeleteAccountResponse struct {
	Message string
}

func (s *Accounts) CreateAccount(request *CreateAccountRequest, response *CreateAccountResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[request.Name]; exists {
		response.Message = "Account already exists."
		return nil
	}

	s.accounts[request.Name] = &Account{Name: request.Name, Balance: request.Balance}
	response.Message = "Account created successfully."
	return nil
}

func (s *Accounts) Deposit(request *DepositRequest, response *DepositResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[request.Name]
	if !exists {
		response.Message = "Account does not exist."
		return nil
	}

	account.Balance += request.Amount
	response.Message = fmt.Sprintf("Deposited %.2f to %s's account. New balance: %.2f", request.Amount, request.Name, account.Balance)
	return nil
}

func (s *Accounts) Withdraw(request *WithdrawRequest, response *WithdrawResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[request.Name]
	if !exists {
		response.Message = "Account does not exist."
		return nil
	}

	if account.Balance < request.Amount {
		response.Message = "Insufficient funds."
		return nil
	}

	account.Balance -= request.Amount
	response.Message = fmt.Sprintf("Withdrew %.2f from %s's account. New balance: %.2f", request.Amount, request.Name, account.Balance)
	return nil
}

func (s *Accounts) ConsultBalance(request *ConsultBalanceRequest, response *ConsultBalanceResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[request.Name]
	if !exists {
		return fmt.Errorf("account does not exist")
	}

	response.Balance = account.Balance
	return nil
}

func (s *Accounts) DeleteAccount(request *DeleteAccountRequest, response *DeleteAccountResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[request.Name]; !exists {
		response.Message = "Account does not exist."
		return nil
	}

	delete(s.accounts, request.Name)
	response.Message = "Account deleted successfully."
	return nil
}

func main() {
	accounts := &Accounts{
		accounts: make(map[string]*Account),
	}

	err := rpc.Register(accounts)
	if err != nil {
		fmt.Println("Error registering RPC:", err)
		return
	}

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running on port 5000...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
