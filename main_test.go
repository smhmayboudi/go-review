// main_test.go
package main_test

import (
	"testing"
)

// تست ساده
func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// تابع کمکی برای تست (در main.go یا اینجا تعریف شود)
func add(a, b int) int {
	return a + b
}
