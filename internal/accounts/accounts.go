package accounts

import (
	"log"

	"github.com/Rhymond/go-money"
)

type Account struct {
	Name string
}

func NewAccount(name string) *Account {
	return &Account{Name: name}
}

func (act *Account) Balance(sourceTransactions []Transaction) *money.Money {
	bal := money.New(0, money.USD)
	entries := make([]Entry, 0, 20)
	for _, t := range sourceTransactions {
		for _, e := range t.Entries {
			if e.Account == act {
				entries = append(entries, e)
			}
		}
	}
	for _, entry := range entries {
		sum, err := bal.Add(entry.Amount)
		if err != nil {
			log.Fatal(err)
		}
		bal = sum

	}
	return bal
}

type Transaction struct {
	Entries []Entry
}

func NewTransaction() *Transaction {
	return &Transaction{Entries: make([]Entry, 0, 4)}
}

func (trans *Transaction) AddEntry(act *Account, amount *money.Money) {
	trans.Entries = append(trans.Entries, Entry{act, amount})
}

func (trans *Transaction) IsValid() bool {
	total := money.New(0, money.USD)
	for _, entry := range trans.Entries {
		sum, err := total.Add(entry.Amount)
		if err != nil {
			log.Fatal(err)
		}
		total = sum
	}
	return total.IsZero()
}

func Transaction2(fromAccount *Account, toAccount *Account, amount *money.Money) *Transaction {
	trans := NewTransaction()
	trans.AddEntry(fromAccount, amount.Negative())
	trans.AddEntry(toAccount, amount)
	return trans
}

type Entry struct {
	Account *Account
	Amount  *money.Money
}
