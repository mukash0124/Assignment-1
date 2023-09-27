package main

type Stock struct {
	observers map[Observer]bool
	name string
	price int
}

func newStock(name string, price int) *Stock  {
	s := &Stock{name: name, price: price}
	s.observers = make(map[Observer]bool)
	return s
}

func (s *Stock) updatePrice(price int) {
	s.price = price
	s.notifyAll()
}

func (s *Stock) register(o Observer) {
    s.observers[o] = true
}

func (s *Stock) unregister(o Observer) {
    delete(s.observers, o) // Or we can just set value to false, so we will not create new key in map, when shareholder will resubscribe. But we will need to store shareholders in map, who unsubscribed.
}


func (s *Stock) notifyAll() {
	for observer := range s.observers {
		observer.update(s.name, s.price)
	}
}