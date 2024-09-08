package accounts

import (
	"log"

	"github.com/Rhymond/go-money"
)

// Account represents a tracking of money.
type Account struct {
	Name string
}

// NewAccount creates a new Account.
func NewAccount(name string) *Account {
	return &Account{Name: name}
}

// getEntries gets all the entries for the Account from the given slice of Transactions.
func (act *Account) getEntries(sourceTransactions []Transaction) []Entry {
	entries := make([]Entry, 0, 20)
	for _, t := range sourceTransactions {
		for _, e := range t.Entries {
			if e.Account == act {
				entries = append(entries, e)
			}
		}
	}
	return entries
}

// Balance calculates the balance of the account, given a list of source transactions.
func (act *Account) Balance(sourceTransactions []Transaction) *money.Money {
	bal := money.New(0, money.USD)
	for _, entry := range act.getEntries(sourceTransactions) {
		sum, err := bal.Add(entry.Amount)
		if err != nil {
			log.Fatal(err)
		}
		bal = sum

	}
	return bal
}

// Transaction represents an atomic transfer of money between accounts.
type Transaction struct {
	Entries []Entry
}

// NewTransaction creates a new Transaction.
func NewTransaction() *Transaction {
	return &Transaction{Entries: make([]Entry, 0, 4)}
}

// AddEntry adds an entry to a transaction.
func (trans *Transaction) AddEntry(act *Account, amount *money.Money) {
	trans.Entries = append(trans.Entries, Entry{act, amount})
}

// IsValid returns a boolean indicating if the transaction is valid.
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

// Transaction2 creates a new Transaction for a transfer between 2 accounts.
func Transaction2(fromAccount *Account, toAccount *Account, amount *money.Money) *Transaction {
	trans := NewTransaction()
	trans.AddEntry(fromAccount, amount.Negative())
	trans.AddEntry(toAccount, amount)
	return trans
}

// Entry represents one entry in a Transaction.
type Entry struct {
	Account *Account
	Amount  *money.Money
}
