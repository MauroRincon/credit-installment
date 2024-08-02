package utils

import "fmt"

// Communication of results
func ResultsCommunication(principal float64, monthlyInterestRate float64, numberOfPayments int) {
	fmt.Println("")
	fmt.Println("La información que ingresaste es: ")
	p.Printf("Valor del crédito: $%.0f\n", principal)
	fmt.Println("Tasa de interés mensual %:", monthlyInterestRate)
	fmt.Println("Cantidad de cuotas:", numberOfPayments)
	fmt.Println("")
}
