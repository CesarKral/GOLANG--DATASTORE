package GAEb

import (
	"time"
)

type User struct {
	Name       string
	Age        int
	Car        string
	Subscribed time.Time
}

type Employee struct {
	Name    string
	Age     int
	Company string
	Country string
}

type Address struct {
	Street     string
	Number     int
	Letter     string
	Floor      int
	PostalCode string
}

type Girlfriend struct {
	Name string
	Age  int
	Eyes string
}

type Wife struct {
	Name string
	Age  int
	Eyes string
}

type Xproj struct {
	Age int
	Car string
}
