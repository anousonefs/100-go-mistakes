package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	/* jsonEmbeddingDemo() */
	/* monotonicClock() */
	/* timeLocation() */
	mapAny()
}

type Event struct {
	ID int
	time.Time
	//Time time.Time
}

// custom Marshaling
func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event // Create an alias to avoid infinite recursion
	return json.Marshal(struct {
		ID        int    `json:"id"`
		Timestamp string `json:"timestamp"`
	}{
		ID:        e.ID,
		Timestamp: e.Time.Format(time.RFC3339),
	})
}

func jsonEmbeddingDemo() error {
	event := Event{
		ID:   1234,
		Time: time.Now(),
	}
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Println(string(b)) //output: "2024-09-03T10:16:39.041199+07:00"
	return nil
}

type Event2 struct {
	Time time.Time
}

func monotonicClock() error {
	t := time.Now()
	event1 := Event2{
		Time: t,
	}
	/* event1 := Event{ */
	/* Time: t.Truncate(0),  */
	/* } */
	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event2
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
	fmt.Println(event1 == event2)               // false
	fmt.Println(event1.Time.Equal(event2.Time)) // true
	return nil
}

func timeLocation() {
	laos, err := time.LoadLocation("Asia/Vientiane")
	if err != nil {
		panic(err)
	}
	t := time.Now().In(laos)
	fmt.Printf("laos time: %v\n", t)

	tu := time.Now().In(laos).UTC()
	fmt.Printf("laos time utc: %v\n\n", tu)

	america, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	t2 := time.Now().In(america)
	fmt.Printf("america time: %v\n", t2)

	t3 := time.Now().In(america).UTC()
	fmt.Printf("america time utc: %v\n", t3)
}

func mapAny() {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("m: %v\n", m)

	fmt.Printf("%T\n", m["key3"])
}

func getMessage() []byte {
	message := map[string]any{
		"key1": "value1",
		"key2": 1234,
		"key3": true,
		"key4": []string{"a", "b", "c"},
	}
	b, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return b
}
