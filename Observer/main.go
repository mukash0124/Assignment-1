package main

func main() {
	stock := newStock("Apple Inc.", 100)

	firstShareholder := &Shareholder{email: "mukash@gmail.com"}
	secondShareholder := &Shareholder{email: "madi@mail.ru"}

	stock.register(firstShareholder)
	stock.register(secondShareholder)

	stock.updatePrice(200)

	stock.unregister(firstShareholder)

	stock.updatePrice(150)
}