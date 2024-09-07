package main

import (
	"fmt"

	"github.com/davidwmartines/accountsim/internal/accounts"
)

func main() {
	fmt.Println("demo!")

	// create accounts
	cash_account := accounts.Account{Name: "cash"}
	income_source_account := accounts.Account{Name: "income source"}

	// slice holding all transactions
	all_transactions := make([]accounts.Transaction, 0, 20)

	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

	// create literal transaction
	trans1 := accounts.Transaction{
		Entries: []accounts.Entry{
			{Account: &income_source_account, Amount: -100},
			{Account: &cash_account, Amount: 100}}}

	all_transactions = append(all_transactions, trans1)
	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

	// append entries post creation
	trans2 := accounts.Transaction{}
	trans2.Entries = append(trans2.Entries, accounts.Entry{Account: &income_source_account, Amount: -250})
	trans2.Entries = append(trans2.Entries, accounts.Entry{Account: &cash_account, Amount: 250})
	all_transactions = append(all_transactions, trans2)
	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

	// using the AddEntry method
	trans3 := accounts.Transaction{}
	trans3.AddEntry(&income_source_account, -10)
	trans3.AddEntry(&cash_account, 10)
	all_transactions = append(all_transactions, trans3)
	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

	// Transaction2 factory function
	trans4 := accounts.Transaction2(&income_source_account, &cash_account, 5)
	all_transactions = append(all_transactions, *trans4)
	print_balance(&cash_account, all_transactions)
	print_balance(&income_source_account, all_transactions)

}

func print_balance(act *accounts.Account, all_transactions []accounts.Transaction) {
	bal := act.Balance(all_transactions)
	fmt.Printf("Account: %s, Balance: %d\n", act.Name, bal)
}
