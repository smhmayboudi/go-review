// package3/main_test.go
package package3_test

import (
	"fmt"
	"testing"

	"github.com/smhmayboudi/go-review/package3"
)

// TestUser_Greet - تست متد Greet
func TestUser_Greet(t *testing.T) {
	tests := []struct {
		name     string
		user     package3.User
		expected string
	}{
		{
			name:     "Normal case",
			user:     package3.User{Name: "Alice", Age: 25},
			expected: "Hello, I'm Alice and I'm 25 years old.",
		},
		{
			name:     "Empty name",
			user:     package3.User{Name: "", Age: 30},
			expected: "Hello, I'm  and I'm 30 years old.",
		},
		{
			name:     "Zero age",
			user:     package3.User{Name: "Bob", Age: 0},
			expected: "Hello, I'm Bob and I'm 0 years old.",
		},
		{
			name:     "Special characters",
			user:     package3.User{Name: "John-Doe", Age: 40},
			expected: "Hello, I'm John-Doe and I'm 40 years old.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.user.Greet()
			if result != tt.expected {
				t.Errorf("Greet() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestUser_FieldAssignment - تست مقداردهی فیلدها
func TestUser_FieldAssignment(t *testing.T) {
	user := package3.User{Name: "Test User", Age: 35}

	if user.Name != "Test User" {
		t.Errorf("Expected Name to be 'Test User', got '%s'", user.Name)
	}

	if user.Age != 35 {
		t.Errorf("Expected Age to be 35, got %d", user.Age)
	}
}

// TestAdd - تست تابع add (با توجه به اینکه private است، باید در همان پکیج تست شود)
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Zeros", 0, 0, 0},
		{"Negative numbers", -1, -2, -3},
		{"Mixed signs", -5, 10, 5},
		{"Large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := package3.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("package3.Add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// BenchmarkUser_Greet - بنچمارک برای متد Greet
func BenchmarkUser_Greet(b *testing.B) {
	user := package3.User{Name: "Benchmark User", Age: 42}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.Greet()
	}
}

// BenchmarkAdd - بنچمارک برای تابع add
func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs() // گزارش تخصیص‌های حافظه

	for i := 0; i < b.N; i++ {
		_ = package3.Add(i, i+1)
	}
}

// BenchmarkAdd_Parallel - بنچمارک موازی
func BenchmarkAdd_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = package3.Add(i, i+1)
			i++
		}
	})
}

// ExampleUser_Greet - مثال documentation
func ExampleUser_Greet() {
	user := package3.User{Name: "Example", Age: 30}
	fmt.Println(user.Greet())
	// Output: Hello, I'm Example and I'm 30 years old.
}

// TestUser_EdgeCases - تست موارد edge case
func TestUser_EdgeCases(t *testing.T) {
	// تست با مقادیر مرزی
	edgeCases := []package3.User{
		{Name: "A", Age: 1},                     // کوتاه‌ترین نام
		{Name: "Very Long Name Here", Age: 150}, // نام طولانی و سن بالا
		{Name: "Test", Age: -1},                 // سن منفی (اگر منطق کسب‌وکار اجازه دهد)
	}

	for _, user := range edgeCases {
		result := user.Greet()
		// فقط بررسی می‌کنیم که خطایی ندهد و خروجی معتبر باشد
		if result == "" {
			t.Errorf("Greet() returned empty string for user %+v", user)
		}
		if len(result) < 10 { // حداقل طول منطقی
			t.Errorf("Greet() returned very short string: %s", result)
		}
	}
}

// TestAdd_EdgeCases - تست موارد edge case برای تابع add
func TestAdd_EdgeCases(t *testing.T) {
	edgeCases := []struct {
		a, b int
	}{
		{0, 0},
		{1<<31 - 1, 1}, // نزدیک به limit int32
		{-1 << 31, 1},  // نزدیک به minimum int32
	}

	for _, tc := range edgeCases {
		result := package3.Add(tc.a, tc.b)
		// بررسی می‌کنیم که نتیجه معکوس باشد
		if package3.Add(tc.b, tc.a) != result {
			t.Errorf("add not commutative for %d and %d", tc.a, tc.b)
		}
	}
}

// TestUser_MethodChaining - تست اینکه متدها propertyهای اصلی را تغییر نمی‌دهند
func TestUser_MethodChaining(t *testing.T) {
	originalUser := package3.User{Name: "Original", Age: 25}

	// فراخوانی متد نباید شیء اصلی را تغییر دهد
	originalUser.Greet()

	if originalUser.Name != "Original" || originalUser.Age != 25 {
		t.Error("Greet method modified the original user")
	}
}
