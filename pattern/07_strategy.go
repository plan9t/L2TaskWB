package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

"Стратегия" - это поведенческий паттерн, который определяет семейство алгоритмов, инкапсулирует каждый из них и делает их
взаимозаменяемыми. Позволяет выбирать алгоритм на лету. Он определяет семейство схожих алгоритмов и помещает каждый из них
в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.


ПЛЮСЫ:
			1. Гибкость и расширяемость: Позволяет добавлять новые стратегии или изменять существующие независимо от контекста.

			2. Избегание условных операторов: Уменьшает зависимость от условных операторов, предоставляя альтернативный способ выбора поведения.

			3. Улучшение читаемости кода: Разделение алгоритмов на отдельные классы делает код более читаемым и поддерживаемым.

МИНУСЫ:
			1. Увеличение числа классов: Каждая стратегия представлена отдельным классом, что может привести к увеличению числа классов в проекте.

			2. Усложнение структуры кода: В небольших проектах или с простыми случаями выбора алгоритма использование стратегии может казаться избыточным.

			3. Затраты на создание объектов: Создание отдельных объектов для каждой стратегии может привести к дополнительным затратам по памяти и производительности.

В примере система для обработки платежей
*/

// Определение интерфейса стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Реализация конкретных стратегий

// Конкретная стратегия 1: Оплата кредитной картой
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. с помощью кредитной карты.", amount)
}

// Конкретная стратегия 2: Оплата через электронный кошелек
type EWalletPayment struct{}

func (e *EWalletPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. через электронный кошелек.", amount)
}

// Контекст, использующий стратегию

// Контекст - структура, которая содержит ссылку на интерфейс стратегии
type PaymentContext struct {
	strategy PaymentStrategy
}

// Метод для установки стратегии в контексте
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Метод для вызова конкретной стратегии
func (p *PaymentContext) PerformPayment(amount float64) string {
	return p.strategy.Pay(amount)
}

func main() {
	// Создаем контекст
	paymentContext := &PaymentContext{}

	// Выбираем стратегию оплаты кредитной картой
	creditCardStrategy := &CreditCardPayment{}
	paymentContext.SetStrategy(creditCardStrategy)
	fmt.Println(paymentContext.PerformPayment(1000.0))

	// Выбираем стратегию оплаты через электронный кошелек
	eWalletStrategy := &EWalletPayment{}
	paymentContext.SetStrategy(eWalletStrategy)
	fmt.Println(paymentContext.PerformPayment(500.0))
}
