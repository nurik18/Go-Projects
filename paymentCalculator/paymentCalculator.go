package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

var paymentHistory []string

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) string {
	payment := fmt.Sprintf("Оплачено %.2f с помощью кредитной карты (%s)", amount, c.CardNumber)
	paymentHistory = append(paymentHistory, payment)
	return payment
}

type PayPal struct {
	email string
}

func (p PayPal) Pay(amount float64) string {
	payment := fmt.Sprintf("Оплачено %.2f через PayPal (%s)", amount, p.email)
	paymentHistory = append(paymentHistory, payment)
	return payment
}

type Bitcoin struct{}

func (b Bitcoin) Pay(amount float64) string {
	payment := fmt.Sprintf("Оплачено %.2f биткоинами", amount)
	paymentHistory = append(paymentHistory, payment)
	return payment
}

type ApplePay struct{}

func (a ApplePay) Pay(amount float64) string {
	payment := fmt.Sprintf("Оплачено %.2f через ApplePay", amount)
	paymentHistory = append(paymentHistory, payment)
	return payment
}

type GooglePay struct{}

func (g GooglePay) Pay(amount float64) string {
	payment := fmt.Sprintf("Оплачено %.2f через GooglePay", amount)
	paymentHistory = append(paymentHistory, payment)
	return payment
}

func ProcessPayment(method PaymentMethod, amount float64) {
	fmt.Println(method.Pay(amount))
}

func showHistory() {
	fmt.Println("\nИстория платежей")
	for i, payment := range paymentHistory {
		fmt.Printf("%d. %s\n", i+1, payment)
	}
}

func main() {
	cc := CreditCard{"1234-5678-9012-3456"}
	pp := PayPal{"nurik@email.com"}
	btc := Bitcoin{}
	ap := ApplePay{}
	gp := GooglePay{}

	ProcessPayment(cc, 100)
	ProcessPayment(pp, 500)
	ProcessPayment(btc, 1000)
	ProcessPayment(ap, 101)
	ProcessPayment(gp, 99.9)

	showHistory()
}
