package main

import "fmt"

func main() {
	storeDemo()
}

type Foo struct {
	ID int
}

type Store struct {
	m map[string]*Foo
}

func (s Store) Put(id string, foo *Foo) {
	s.m[id] = foo
}

// -----------

type LargeStruct struct {
	foo string
}

// copy
func updateMapValue(mapValue map[string]LargeStruct, id string) {
	value := mapValue[id]
	value.foo = "bar"
	mapValue[id] = value
}

// pointer
func updateMapPointer(mapPointer map[string]*LargeStruct, id string) {
	mapPointer[id].foo = "bar"
}

// ---------------

type Customer struct {
	ID      string
	Balance float64
}

type StoreCustomer struct {
	m map[string]*Customer
}

func (s *StoreCustomer) storeCustomers(customers []Customer) {
	if s.m == nil {
		s.m = make(map[string]*Customer)
	}
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		cust := customer
		s.m[customer.ID] = &cust
		/* There's no longer this issue from Go 1.22: */
		/* s.m[customer.ID] = &customer */
	}
}

func storeDemo() {
	var s StoreCustomer
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	for _, v := range s.m {
		fmt.Printf("id: %v, balance: %v\n", v.ID, v.Balance)
	}
}
