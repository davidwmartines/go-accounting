package main

import (
	"fmt"

	"github.com/davidwmartines/accountsim/internal/accounts"
)

func main() {
	fmt.Println("demo!")

	cash_account := accounts.Account{Name: "cash"}

	income_source_account := accounts.Account{Name: "income source"}

	all_transactions := make([]accounts.Transaction, 0, 20)

	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

	trans1 := accounts.Transaction{
		Entries: []accounts.Entry{
			{Account: &income_source_account, Amount: -100},
			{Account: &cash_account, Amount: 100}}}

	all_transactions = append(all_transactions, trans1)

	trans2 := accounts.Transaction{
		Entries: []accounts.Entry{
			{Account: &income_source_account, Amount: -250},
			{Account: &cash_account, Amount: 250}}}

	all_transactions = append(all_transactions, trans2)

	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

}

func print_balance(act *accounts.Account, all_transactions []accounts.Transaction) {
	bal := act.Balance(all_transactions)
	fmt.Printf("Account: %s balance: %d\n", act.Name, bal)
}
