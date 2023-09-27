package main

type Account struct {
	discount Discount
	money int
}

func createAccount() *Account {
	return &Account{
		discount: &Base{},
		money: 100,
	}
}

func (a *Account) setStrategy(d Discount) {
	a.discount = d
}

func (a *Account) checkout() {
	a.discount.purchase(a.money)
}