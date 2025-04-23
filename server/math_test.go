package server

import (
	"testing"
)

// Test generated using Keploy
func TestFibonacci_ValidAndNegativeInput_606(t *testing.T) {
	result, err := fibonacci(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{0, 1, 1, 2, 3}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d but got %d at index %d", v, result[i], i)
		}
	}

	_, err = fibonacci(-3)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestIsHappyNumber_ValidInputs_1616(t *testing.T) {
	if !isHappyNumber(19) {
		t.Errorf("Expected true but got false for happy number 19")
	}

	if isHappyNumber(4) {
		t.Errorf("Expected false but got true for non-happy number 4")
	}
}

// Test generated using Keploy

func TestIsAnagram_ValidInputs_1313(t *testing.T) {
	if !isAnagram("listen", "silent") {
		t.Errorf("Expected true but got false for anagrams 'listen' and 'silent'")
	}

	if isAnagram("hello", "world") {
		t.Errorf("Expected false but got true for non-anagrams 'hello' and 'world'")
	}
}

// Test generated using Keploy

func TestFactorial_ValidZeroAndNegativeInput_505(t *testing.T) {
	result, err := factorial(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 120
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	result, err = factorial(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected = 1
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	_, err = factorial(-3)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestPrimeFactors_ValidInputs_1010(t *testing.T) {
	result := primeFactors(28)
	expected := []int{2, 2, 7}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d but got %d at index %d", v, result[i], i)
		}
	}

	result = primeFactors(1)
	expected = []int{}
	if len(result) != len(expected) {
		t.Errorf("Expected empty slice but got %v", result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_ValidInputs_1111(t *testing.T) {
	if !isPalindrome("racecar") {
		t.Errorf("Expected true but got false for palindrome 'racecar'")
	}

	if isPalindrome("hello") {
		t.Errorf("Expected false but got true for non-palindrome 'hello'")
	}
}

// Test generated using Keploy

func TestIsPrime_PrimeAndNonPrimeNumbers_909(t *testing.T) {
	if !isPrime(7) {
		t.Errorf("Expected true but got false for prime number 7")
	}

	if isPrime(4) {
		t.Errorf("Expected false but got true for non-prime number 4")
	}

	if isPrime(1) {
		t.Errorf("Expected false but got true for number 1")
	}

	if isPrime(-5) {
		t.Errorf("Expected false but got true for negative number -5")
	}
}

// Test generated using Keploy

func TestIsPerfectNumber_ValidInputs_1515(t *testing.T) {
	if !isPerfectNumber(28) {
		t.Errorf("Expected true but got false for perfect number 28")
	}

	if isPerfectNumber(10) {
		t.Errorf("Expected false but got true for non-perfect number 10")
	}
}

// Test generated using Keploy

func TestIsFibonacci_ValidInputs_1717(t *testing.T) {
	if !isFibonacci(8) {
		t.Errorf("Expected true but got false for Fibonacci number 8")
	}

	if isFibonacci(10) {
		t.Errorf("Expected false but got true for non-Fibonacci number 10")
	}
}

// Test generated using Keploy

func TestReverseString_ValidInputs_1212(t *testing.T) {
	result := reverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	result = reverseString("")
	expected = ""
	if result != expected {
		t.Errorf("Expected empty string but got %s", result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_ValidInputs_1818(t *testing.T) {
	if !isPerfectSquare(16) {
		t.Errorf("Expected true but got false for perfect square 16")
	}

	if isPerfectSquare(15) {
		t.Errorf("Expected false but got true for non-perfect square 15")
	}
}

// Test generated using Keploy

func TestDivide_ValidAndZeroDivision_101(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	_, err = divide(10, 0)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestModulus_ValidAndZeroDivision_202(t *testing.T) {
	result, err := modulus(10, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 1
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	_, err = modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestSquareRoot_ValidAndNegativeInput_404(t *testing.T) {
	result, err := squareRoot(16)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 4
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	_, err = squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestMultiply_ValidInputs_789(t *testing.T) {
	result := multiply(3, 7)
	expected := 21
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSum_ValidInputs_123(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_ValidInputs_456(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestPower_ValidInputs_303(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestLCM_ValidInputs_808(t *testing.T) {
	result := lcm(12, 15)
	expected := 60
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	result = lcm(0, 15)
	expected = 0
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy
