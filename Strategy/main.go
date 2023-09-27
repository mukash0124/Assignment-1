package main

func main() {
	firstAccount := createAccount()

	firstAccount.checkout()

	premium := &Premium{}

	firstAccount.setStrategy(premium)

	firstAccount.checkout()

	secondAccount := &Account{discount: premium, money: 200}

	secondAccount.checkout()
}