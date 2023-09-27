package main

import "fmt"

type Shareholder struct {
	email string
}

func (s *Shareholder) update(stockName string, stockPrice int) {
	fmt.Printf("Sending email to shareholder %s for change in price of %s stock. New price is %d\n", s.email, stockName, stockPrice)
}