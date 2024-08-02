package utils

import (
	"fmt"
	"strings"
)

func FeesData(principal float64, numberOfPayments int, convertMonthlyInterestRate float64, monthlyPayment float64) {
	// Print headers
	headers := []string{"Cuota N°", "Cuota mensual", "Pago a capital", "Pago intereses", "Saldo"}
	fmt.Printf("%-10s | %-10s | %-10s | %-10s | %-10s\n", headers[0], headers[1], headers[2], headers[3], headers[4])

	// Print dividing line
	fmt.Println(strings.Repeat("-", 74))

	remainingBalance := principal

	for i := 1; i <= numberOfPayments; i++ {
		interestPayment := remainingBalance * convertMonthlyInterestRate
		principalPayment := monthlyPayment - interestPayment
		remainingBalance -= principalPayment
		p.Printf("%-10v $%-15.0f $%-15.0f $%-15.0f $%-15.0f\n", i, monthlyPayment, principalPayment, interestPayment, remainingBalance)
	}
}
