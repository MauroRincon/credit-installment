package utils

import (
	"math"
)

func CreditCalculation(principal float64, monthlyInterestRate float64, numberOfPayments int) (float64, float64) {
	convertMonthlyInterestRate := monthlyInterestRate / 100

	var annualInterestRate float64 = (math.Pow(1+convertMonthlyInterestRate, 12) - 1) * 100

	//annualInterestRate := convertMmonthlyInterestRate * 12 // Convert annual interest rate to monthly
	monthlyPayment := calculatePayment(principal, convertMonthlyInterestRate, numberOfPayments)

	return monthlyPayment, annualInterestRate
}

// calculatePayment calculates the periodic payment of an annuity.
func calculatePayment(pv, r float64, n int) float64 {
	payment := (pv * r * math.Pow(1+r, float64(n))) / (math.Pow(1+r, float64(n)) - 1)
	return payment
}
