package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"golang.org/x/text/message"
)

var p *message.Printer = message.NewPrinter(message.MatchLanguage("en"))

func main() {
	var principal float64
	fmt.Print("Ingresa el valor del crédito (préstamo): ")
	fmt.Scan(&principal)

	var monthlyInterestRate float64
	fmt.Print("Ingresa la tasa de interés mensual: ")
	fmt.Scan(&monthlyInterestRate)

	convertMonthlyInterestRate := monthlyInterestRate / 100

	var annualInterestRate float64 = (math.Pow(1+convertMonthlyInterestRate, 12) - 1) * 100

	var numberOfPayments int
	fmt.Print("Ingresa la cantidad de cuotas: ")
	fmt.Scan(&numberOfPayments)

	resultsCommunication(principal, monthlyInterestRate, numberOfPayments)

	//annualInterestRate := convertMmonthlyInterestRate * 12 // Convert annual interest rate to monthly
	monthlyPayment := calculatePayment(principal, convertMonthlyInterestRate, numberOfPayments)

	fmt.Println("RESULTADOS")
	fmt.Printf("Tasa de interés efectiva anual %%: %.2f\n", annualInterestRate)
	p.Printf("El pago mensual sería: $%.0f\n", monthlyPayment)
	fmt.Println("")

	feesData(principal, numberOfPayments, convertMonthlyInterestRate, monthlyPayment)

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
