package domain

import (
	"time"
)

type AccountID int

type Money int

type Account struct {
	ID              AccountID
	BaselineBalance Money
	ActivityWindow  ActivityWindow
}

func (a *Account) CalculateBalance() Money {
	return AddMoney(
		a.BaselineBalance,
		a.ActivityWindow.CalculateBalance(a.ID))
}

func (a *Account) Withdraw(money Money, targetAccountID AccountID) bool {
	if !a.mayWithdraw(money) {
		return false
	}

	withdrawal := NewActivity(a.ID, targetAccountID, time.Now(), money)
	a.ActivityWindow.AddActivity(withdrawal)
	return true
}

func (a *Account) mayWithdraw(money Money) bool {
	return AddMoney(
		a.CalculateBalance(),
		money.Negate()).IsPositive()
}

func AddMoney(balance Money, amount Money) Money {
	return balance + amount
}

func (m Money) Negate() Money {
	return -m
}

func (a *Account) Deposit(money Money, sourceAccountID AccountID) bool {
	deposit := NewActivity(a.ID, sourceAccountID, time.Now(), money)
	a.ActivityWindow.AddActivity(deposit)
	return true
}

func (m Money) IsPositive() bool {
	return m > 0
}
