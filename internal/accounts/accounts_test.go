package accounts

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func Test_Transaction2_Creates2Entries(t *testing.T) {
	account1 := NewAccount("1")
	account2 := NewAccount("2")
	trans := Transaction2(account1, account2, money.New(100, money.USD))
	assert.True(t, len(trans.Entries) == 2, "transaction should have 2 entries")
}

func Test_Transasction_IsValid_True(t *testing.T) {
	account1 := NewAccount("1")
	account2 := NewAccount("2")

	trans := NewTransaction()

	trans.AddEntry(account1, money.New(100, money.USD).Negative())
	trans.AddEntry(account2, money.New(100, money.USD))

	assert.True(t, trans.IsValid(), "transaction should be valid")
}

func Test_Transasction_IsValid_False(t *testing.T) {
	account1 := NewAccount("1")
	account2 := NewAccount("2")

	trans := NewTransaction()

	trans.AddEntry(account1, money.New(100, money.USD).Negative())
	trans.AddEntry(account2, money.New(101, money.USD))

	assert.False(t, trans.IsValid(), "transaction should not be valid")
}

func Test_Account_Balance_NoTransactions(t *testing.T) {
	account1 := NewAccount("1")
	balance := account1.Balance([]Transaction{})
	assert.True(t, balance.IsZero(), "balance should be 0")
}

func Test_Account_Balance(t *testing.T) {
	account1 := NewAccount("1")
	account2 := NewAccount("2")

	t1 := Transaction2(account1, account2, money.New(100, money.USD))

	bal1 := account1.Balance([]Transaction{*t1})
	bal2 := account2.Balance([]Transaction{*t1})
	assert.Equal(t, money.New(100, money.USD).Negative(), bal1, "Balance of account1 should be -100.")
	assert.Equal(t, money.New(100, money.USD), bal2, "Balance of account2 should be 100.")
}
