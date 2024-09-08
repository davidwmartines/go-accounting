package main

import (
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/davidwmartines/accountsim/internal/accounts"
)

func main() {
	fmt.Println("demo!")

	// create accounts
	cashAccount := accounts.NewAccount("cash")
	incomeSourceAccount := accounts.NewAccount("income source")

	// slice holding all transactions
	all_transactions := make([]accounts.Transaction, 0, 20)

	printBalance(cashAccount, all_transactions)
	printBalance(incomeSourceAccount, all_transactions)

	// create literal transaction
	trans1 := accounts.Transaction{
		Entries: []accounts.Entry{
			{Account: incomeSourceAccount, Amount: money.NewFromFloat(-100, money.USD).Negative()},
			{Account: cashAccount, Amount: money.NewFromFloat(100, money.USD)}}}

	checkTransaction(&trans1)
	all_transactions = append(all_transactions, trans1)
	printBalance(cashAccount, all_transactions)
	printBalance(incomeSourceAccount, all_transactions)

	// append entries post creation
	trans2 := accounts.NewTransaction()
	trans2.Entries = append(trans2.Entries, accounts.Entry{Account: incomeSourceAccount, Amount: money.NewFromFloat(250, money.USD).Negative()})
	trans2.Entries = append(trans2.Entries, accounts.Entry{Account: cashAccount, Amount: money.NewFromFloat(250, money.USD)})
	checkTransaction(trans2)
	all_transactions = append(all_transactions, *trans2)
	printBalance(cashAccount, all_transactions)
	printBalance(incomeSourceAccount, all_transactions)

	// using the AddEntry method
	trans3 := accounts.NewTransaction()
	trans3.AddEntry(incomeSourceAccount, money.NewFromFloat(10, money.USD).Negative())
	trans3.AddEntry(cashAccount, money.NewFromFloat(10, money.USD))
	checkTransaction(trans3)
	all_transactions = append(all_transactions, *trans3)
	printBalance(cashAccount, all_transactions)
	printBalance(incomeSourceAccount, all_transactions)

	// Transaction2 factory function
	trans4 := accounts.Transaction2(incomeSourceAccount, cashAccount, money.NewFromFloat(5, money.USD))
	checkTransaction(trans4)
	all_transactions = append(all_transactions, *trans4)
	printBalance(cashAccount, all_transactions)
	printBalance(incomeSourceAccount, all_transactions)

	// invalid transaction
	trans5 := accounts.NewTransaction()
	trans5.AddEntry(incomeSourceAccount, money.NewFromFloat(100, money.USD).Negative())
	trans5.AddEntry(cashAccount, money.NewFromFloat(101, money.USD))
	checkTransaction(trans5)

}

func checkTransaction(trans *accounts.Transaction) {
	fmt.Printf("Transaction is valid: %t\n", trans.IsValid())
}

func printBalance(act *accounts.Account, all_transactions []accounts.Transaction) {
	bal := act.Balance(all_transactions)
	fmt.Printf("Account: %s, Balance: %s\n", act.Name, bal.Display())
}
