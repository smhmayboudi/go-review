package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello I am %s, and %d years old.", u.Name, u.Age)
}

func main() {
	user := User{Name: "Hossein Mayboudi", Age: 45}
	fmt.Println(user.Greet())
}
