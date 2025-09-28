package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello I am %s, and %d years old.", u.Name, u.Age)
}

func getUser() *User {
	// user یک متغیر محلی است، اما چون آدرس آن برگردانده می‌شود (فرار می‌کند)...
	user := User{Name: "Hossein Mayboudi", Age: 45}
	// ... کامپایلر Go به طور خودکار آن را در هیپ تخصیص می‌دهد.
	// بنابراین پس از بازگشت از تابع، شیء از بین نمی‌رود.
	return &user // برگرداندن آدرس (pointer)
}

func main() {
	user := User{Name: "Hossein Mayboudi", Age: 45}
	fmt.Println(user.Greet())
	userPtr := getUser() // شیء User در هیپ زنده می‌ماند.
	fmt.Println(userPtr.Greet())
}

// Garbage Collector در نهایت این شیء هیپ را وقتی دیگر استفاده نشد، پاک می‌کند.
