package accounts

type Account struct {
	Name string
}

func NewAccount(name string) *Account {
	return &Account{Name: name}
}

func (act *Account) Balance(source_transactions []Transaction) int {
	var bal int
	entries := make([]Entry, 0, 20)
	for _, t := range source_transactions {
		for _, e := range t.Entries {
			if e.Account == act {
				entries = append(entries, e)
			}
		}
	}
	for _, entry := range entries {
		bal += entry.Amount
	}
	return bal
}

type Transaction struct {
	Entries []Entry
}

func NewTransaction() *Transaction {
	return &Transaction{Entries: make([]Entry, 0, 4)}
}

func (trans *Transaction) AddEntry(act *Account, amount int) {
	trans.Entries = append(trans.Entries, Entry{act, amount})
}

func Transaction2(fromAccount *Account, toAccount *Account, amount int) *Transaction {
	trans := NewTransaction()
	trans.AddEntry(fromAccount, amount*-1)
	trans.AddEntry(toAccount, amount)
	return trans
}

type Entry struct {
	Account *Account
	Amount  int
}
