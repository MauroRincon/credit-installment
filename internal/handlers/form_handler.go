package handlers

import (
	"credit-installment/internal/utils"
	"html/template"
	"net/http"
	"strconv"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	// Obtener el valor del input como string
	principalStr := r.FormValue("inputValue")
	monthlyInterestRateStr := r.FormValue("inputInterestRate")
	numberOfPaymentsStr := r.FormValue("inputNumberPayments")

	// Convertir el valor de string a float64 o int
	principal, err := strconv.ParseFloat(principalStr, 64)
	if err != nil {
		http.Error(w, "Invalid input value", http.StatusBadRequest)
		return
	}

	monthlyInterestRate, err := strconv.ParseFloat(monthlyInterestRateStr, 64)
	if err != nil {
		http.Error(w, "Invalid input value", http.StatusBadRequest)
		return
	}

	numberOfPayments, err := strconv.Atoi(numberOfPaymentsStr)
	if err != nil {
		http.Error(w, "Invalid input value", http.StatusBadRequest)
		return
	}

	monthlyPayment, annualInterestRate := utils.CreditCalculation(principal, monthlyInterestRate, numberOfPayments)

	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"calculatePayment":   monthlyPayment,
		"annualInterestRate": annualInterestRate,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
