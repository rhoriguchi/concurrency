package controller

import (
	"errors"
	"pkgArch/entity"
	"time"
)

type UserController interface {
	Notify(usr entity.User, msg string)
}

type AccountController struct {
	UserCtl UserController
}

func (a *AccountController) Transfer(from, to *entity.Account, amount float64, when time.Time) error {
	from.Mu.RLock()

	// nail this down for write lock promotion
	currentFunds := from.Balance

	// verify 'when' is not too far in the past
	if time.Now().Add(-time.Second).After(when) {
		return errors.New("can't transfer funds in the past")
	}

	// refuse to block all operations if 'when' is too far out
	if when.Sub(time.Now()).Seconds() > 10 {
		return errors.New("scheduled time is too far out")
	}

	// verify funds are sufficient
	if from.Balance < amount {
		return errors.New("insufficient funds")
	}

	if amount < 0 {
		return errors.New("dude, someone trys to steal money")
	}

	from.Mu.RUnlock()

	// Other goroutines might obtain the lock in the meantime! That's fine as long as we check the balance again below.

	// FIXME potential deadlock
	from.Mu.Lock()
	to.Mu.Lock()
	defer from.Mu.Unlock()
	defer to.Mu.Unlock()

	if from.Balance != currentFunds {
		return errors.New("invalid transaction: balance changed while verification in progress")
	}

	// wait until we reach scheduled time
	if when.After(time.Now()) { // you never know how long we wait for the mutex
		time.Sleep(when.Sub(time.Now()))
	}

	// all good: proceed with transfer
	from.Balance -= amount
	to.Balance += amount

	a.UserCtl.Notify(*from.Owner, "Money has been taken from your account. Glad to be of service.")
	a.UserCtl.Notify(*to.Owner, "What a wonderful day, you received money!")

	return nil
}
