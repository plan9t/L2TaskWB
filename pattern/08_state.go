package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern


*/

// Определение интерфейса состояния
type State interface {
	TransitionToNextState() State
	GetColor() string
}

// Реализация конкретных состояний

// Для красного
type RedState struct{}

func (r *RedState) TransitionToNextState() State {
	return &YellowState{}
}

func (r *RedState) GetColor() string {
	return "Красный"
}

// Для жёлтого
type YellowState struct{}

func (y *YellowState) TransitionToNextState() State {
	return &GreenState{}
}

func (y *YellowState) GetColor() string {
	return "Жёлтый"
}

// Для зелёного
type GreenState struct{}

func (g *GreenState) TransitionToNextState() State {
	return &RedState{}
}

func (g *GreenState) GetColor() string {
	return "Зелёный"
}

// Реализация контекста, здесь это светофор
type TrafficLight struct {
	state State
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{
		state: &RedState{},
	}
}

func (t *TrafficLight) ChangeState() {
	t.state = t.state.TransitionToNextState()
	time.Sleep(time.Millisecond * 1000)
}

func (t *TrafficLight) ShowColor() {
	color := t.state.GetColor()
	fmt.Println("Светофор", color)
}

// Step 4: Использование паттерна
func main() {
	// Создаем светофор (экземпляр)
	trafficLight := NewTrafficLight()

	// Имитация работы светофора в течение трех шагов
	for i := 0; i < 3; i++ {
		trafficLight.ShowColor()
		trafficLight.ChangeState()
	}
}
