package main

import (
	"fmt"
	"math"

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

	convertMmonthlyInterestRate := monthlyInterestRate / 100

	var numberOfPayments int
	fmt.Print("Ingresa la cantidad de cuotas: ")
	fmt.Scan(&numberOfPayments)

	fmt.Println("La información que ingresaste es: ")
	p.Printf("Valor del crédito: $%.0f\n", principal)
	fmt.Println("Valor del crédito:", monthlyInterestRate)
	fmt.Println("Valor del crédito:", numberOfPayments)
	fmt.Println("")

	//annualInterestRate := convertMmonthlyInterestRate * 12 // Convert annual interest rate to monthly
	monthlyPayment := calculatePayment(principal, convertMmonthlyInterestRate, numberOfPayments)
	p.Printf("El pago mensual sería: $%.0f\n", monthlyPayment)
}

// calculatePayment calculates the periodic payment of an annuity.
func calculatePayment(pv, r float64, n int) float64 {
	payment := (pv * r * math.Pow(1+r, float64(n))) / (math.Pow(1+r, float64(n)) - 1)
	return payment
}
