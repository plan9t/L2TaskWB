package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Паттерн "Команда" (Command) является поведенческим паттерном проектирования,
который инкапсулирует запрос в виде объекта, позволяя параметризовать клиентов с различными запросами,
организовывать запросы в очередь, а также поддерживать отмену операций.


ПЛЛЮСЫ:
			1. Разделение отправителя и получателя: Паттерн "Команда" позволяет разделить объект, инициирующий запрос (отправитель) от объекта,
			который знает, как выполнить этот запрос (получатель). Это уменьшает связанность между отправителем и получателем.

			2. Отмена операций: Паттерн обеспечивает легкость в реализации механизма отмены операций. Каждая команда может иметь
			метод отмены (undo), что полезно, когда необходимо отменять выполненные операции.

			3. Гибкость в добавлении новых команд: Вы можете легко добавлять новые команды, расширяя классы команд,
			не изменяя клиентский код. Это поддерживает открытость/закрытость принцип из SOLID.

			4. Возможность реализации очереди команд: Команды могут быть сохранены в стеке или очереди, что позволяет реализовать
			различные сценарии, такие как отмена и повтор действий.


МИНУСЫ:
			1. Увеличение числа классов (структур в случае с Go): Каждая команда требует создания отдельного класса, что может привести к
			увеличению числа классов в программе. В небольших проектах это может быть излишним.

			2. Сложность отслеживания изменений в состоянии: Когда команды взаимодействуют с объектами, изменяя их состояние,
			это может усложнить отслеживание изменений в системе, особенно если команды динамически изменяются во времени.

			3. Сложность в поддержке отмены операций: Реализация отмены операций может быть сложной, особенно если требуется учет всех
			изменений, связанных с выполнением команд.

			4. Дополнительные затраты по памяти: Использование очереди или стека для управления командами может привести к дополнительным
			затратам по памяти, особенно при хранении большого числа команд.

В примере ситуация, где у нас есть светильник, и мы хотим реализовать функциональность включения и выключения света с использованием команд.

*/

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретная команда для включения света
type LightOnCommand struct {
	light *Light
}

// Execute реализация метода Execute для включения света
func (c *LightOnCommand) Execute() {
	c.light.turnOn()
}

// Конкретная команда для выключения света
type LightOffCommand struct {
	light *Light
}

// Execute реализация метода Execute для выключения света
func (c *LightOffCommand) Execute() {
	c.light.turnOff()
}

// Получатель команды - объект, который выполняет действие (светильник)
type Light struct {
	isOn bool
}

// Методы для включения и выключения света
func (l *Light) turnOn() {
	l.isOn = true
	fmt.Println("Свет включен")
}

func (l *Light) turnOff() {
	l.isOn = false
	fmt.Println("Свет выключен")
}

// Инвокер - объкт, который вызывает команды
type RemoteControl struct {
	command Command
}

// SetCommand метод для установки команды в инвокер
func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

// PressButton метод для выполнения команды
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	// Создаем экземпляры светильника и команд
	light := &Light{}
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// Создаем экземпляр пульта управления
	remote := &RemoteControl{}

	// Настройка команды на включение света и выполнение
	remote.SetCommand(lightOn)
	remote.PressButton()

	// Настройка команды на выключение света и выполнение
	remote.SetCommand(lightOff)
	remote.PressButton()
}