package handlers

import (
	"credit-installment/internal/utils"
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type PaymentRow struct {
	QuotaNo          int
	CapitalPayment   string
	InterestsPayment string
	Balance          string
}

type TableData struct {
	Rows []PaymentRow
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
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
	// Get input value as a string
	principalStr := r.FormValue("inputValue")
	monthlyInterestRateStr := r.FormValue("inputInterestRate")
	numberOfPaymentsStr := r.FormValue("inputNumberPayments")

	// Convert to float or int
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
	p := message.NewPrinter(language.English)
	formattedmonthlyPayment := p.Sprintf("$%.0f", monthlyPayment)
	formattedAnnualRate := strconv.FormatFloat(annualInterestRate, 'f', 2, 64)

	// Table with monthly rates
	convertMonthlyInterestRate := monthlyInterestRate / 100

	remainingBalance := principal
	totalPaid := monthlyPayment * float64(numberOfPayments)
	formattedTotalPaid := p.Sprintf("$%.0f", totalPaid)
	var rows []PaymentRow

	for i := 1; i <= numberOfPayments; i++ {
		interestPayment := remainingBalance * convertMonthlyInterestRate
		principalPayment := monthlyPayment - interestPayment
		remainingBalance -= principalPayment

		// Add values to the rows
		rows = append(rows, PaymentRow{
			QuotaNo:          i,
			CapitalPayment:   p.Sprintf("$%.0f", principalPayment),
			InterestsPayment: p.Sprintf("$%.0f", interestPayment),
			Balance:          p.Sprintf("$%.0f", remainingBalance),
		})

		totalPaid += monthlyPayment
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"calculatePayment":   formattedmonthlyPayment,
		"annualInterestRate": formattedAnnualRate,
		"totalPaid":          formattedTotalPaid,
		"rows":               rows,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
