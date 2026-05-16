package accounts

type Account struct {
	ID int
	Owner string
	balance float64
}

func (a *Account) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}

	a.balance += amount
	return true
}

func (a *Account) Balance() float64 {
	return a.balance
}