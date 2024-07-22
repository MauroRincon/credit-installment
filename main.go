package main

import (
	"fmt"
	"math"
)

// calculatePayment calculates the periodic payment of an annuity.
func calculatePayment(pv, r float64, n int) float64 {
	payment := (pv * r * math.Pow(1+r, float64(n))) / (math.Pow(1+r, float64(n)) - 1)
	return payment
}

func main() {
	principal := 100000.0      // Present value or loan amount
	annualInterestRate := 0.05 // Annual interest rate
	numberOfPayments := 360    // Number of payments (e.g., 30 years of monthly payments)

	monthlyInterestRate := annualInterestRate / 12 // Convert annual interest rate to monthly

	monthlyPayment := calculatePayment(principal, monthlyInterestRate, numberOfPayments)
	fmt.Printf("The monthly payment is: %.2f\n", monthlyPayment)
}
