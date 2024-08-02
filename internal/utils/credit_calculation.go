package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"golang.org/x/text/message"
)

var p *message.Printer = message.NewPrinter(message.MatchLanguage("en"))

func CreditCalculation() {
	var principal float64
	var monthlyInterestRate float64

	convertMonthlyInterestRate := monthlyInterestRate / 100

	var annualInterestRate float64 = (math.Pow(1+convertMonthlyInterestRate, 12) - 1) * 100

	var numberOfPayments int

	ResultsCommunication(principal, monthlyInterestRate, numberOfPayments)

	//annualInterestRate := convertMmonthlyInterestRate * 12 // Convert annual interest rate to monthly
	monthlyPayment := calculatePayment(principal, convertMonthlyInterestRate, numberOfPayments)

	fmt.Println("RESULTADOS")
	fmt.Printf("Tasa de interés efectiva anual %%: %.2f\n", annualInterestRate)
	p.Printf("El pago mensual sería: $%.0f\n", monthlyPayment)
	fmt.Println("")

	FeesData(principal, numberOfPayments, convertMonthlyInterestRate, monthlyPayment)

	// Pause execution until the user presses Enter
	fmt.Println("")
	fmt.Println("Presione 'Enter' para salir...")
	var exit int
	fmt.Scanln(&exit)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// calculatePayment calculates the periodic payment of an annuity.
func calculatePayment(pv, r float64, n int) float64 {
	payment := (pv * r * math.Pow(1+r, float64(n))) / (math.Pow(1+r, float64(n)) - 1)
	return payment
}
