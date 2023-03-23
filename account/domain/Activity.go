package domain

import (
	"time"
)

type Activity struct {
	SourceAccountID AccountID
	TargetAccountID AccountID
	Timestamp       time.Time
	Amount          Money
}

func NewActivity(sourceAccountID, targetAccountID AccountID, timestamp time.Time, amount Money) Activity {
	return Activity{
		SourceAccountID: sourceAccountID,
		TargetAccountID: targetAccountID,
		Timestamp:       timestamp,
		Amount:          amount,
	}
}

type ActivityWindow struct {
	Activities []Activity
}

func (w *ActivityWindow) CalculateBalance(accountID AccountID) Money {
	var balance Money
	for _, activity := range w.Activities {
		if activity.SourceAccountID == accountID {
			balance -= activity.Amount
		} else if activity.TargetAccountID == accountID {
			balance += activity.Amount
		}
	}
	return balance
}

func (w *ActivityWindow) AddActivity(activity Activity) {
	w.Activities = append(w.Activities, activity)
}
