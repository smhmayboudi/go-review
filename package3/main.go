// package3/main.go
package package3

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello, I'm %s and I'm %d years old.", u.Name, u.Age)
}

func Add(a int, b int) int {
	return a + b
}
