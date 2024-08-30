package main

import (
	"fmt"
	"sync"
)

/* func (w *watcher) Watch(ctx context.Context, key string, */
/*     opts ...OpOption) WatchChan { */
/*     // ... */
/* ctxKey := fmt.Sprintf("%v", ctx) ❶ // ... */
/* wgs := w.streams[ctxKey] */
/* // ... */

/*
	where one goroutine was updating one of the context values,
	whereas another was executing Watch,
	hence reading all the values in this context.
	This led to a data race.
*/

//--------------- deadlock -----------------

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

// deadlock!
func (c *Customer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c) // call String() method -> Lock mutex
	}
	c.age = age
	return nil
}

func (c *Customer) UpdateAgeV2(age int) error {
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.age = age
	return nil
}

func (c *Customer) UpdateAgeV3(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if age < 0 {
		return fmt.Errorf("age should be positive for customer id %s", c.id) // don't call String() method
	}
	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func main() {
	customer := &Customer{id: "12345", age: 25}
	fmt.Println(customer.String())

	err := customer.UpdateAge(30)
	if err != nil {
		fmt.Println("Error updating age:", err)
	} else {
		fmt.Println("Customer's age updated successfully")
	}

	fmt.Println(customer.String())

	err = customer.UpdateAgeV3(-5)
	if err != nil {
		fmt.Println("Error updating age:", err)
	}

	fmt.Println(customer.String())
}
