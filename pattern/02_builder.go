package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Строитель - порождающий паттер. Он позволяет создавать сложные объекты используя "шаги".
На каждом шаге производится какая-то часть общего (финального) объекта (или продукта, в моей реализации это компьютер).
Выполняя шаги по очереди мы формируем конечный объект.

Строитель даёт возможность использовать один и тот же код строительства объекта для получения РАЗНЫХ представлений этого объекта.
Какой-то шаг можно пропустить, а можно наоборот добавить, если это потребуется.
Для каждого такого строителя применяется общий интерфейс простройки.

ПЛЮСЫ: 1. Общий продукт создаётся пошагово.
	   2. Строитель позволяет использовать один и тот же код для создания отличных друг от друга объектов.
	   3. Изолирует сложный код сборки и его основную бизнес-логику.

МИНУСЫ: 1. Усложняет код программы из-за введения дополнительных структур, интерфесов и т.д
		2. Клиент привязывается к конкретному объекту строителя, а в интерфейсе может не быть какого-то метода, тогда его нужно будет добавлять вручную.

Реализация на примере завода по сборке компьютеров для РАЗНЫХ брендов. Мы будем менять бренд компьютера и его комплектующие
не изменяя саму логику сборки.





P.S: В реализации у меня всё находится в одном пакете согласно структуре. Правильнее было бы разделить логику
магазина, банка и пользователя в отдельные пакеты.
*/

// Типы сборщиков
const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

// Интерфейс сборщика
type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

// Структура компьютера
type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf("%s, Core: [%d], Mem: [%d], Graphic: [%d], Monitor: [%d]\n", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}

// Для бренда ASUS
type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "ASUS"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

// Для  бренда HP
type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 6
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "HP"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 2
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

// Структура самого завода по производству компьютеров
type Factory struct {
	Collector Collector
}

// Функция иниализации нового завода
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

// Функция, меняющее поведение завода в зависимости от типа сборщика (для асус или хп)
func (f *Factory) SetCollecor(collector Collector) {
	f.Collector = collector
}

// Функция, создающая компьютер. СТРОИТЕЛЬ
func (f *Factory) CreateComputer() Computer {
	f.Collector.SetCore()        // ШАГ
	f.Collector.SetBrand()       // ШАГ
	f.Collector.SetMemory()      // ШАГ
	f.Collector.SetMonitor()     // ШАГ
	f.Collector.SetGraphicCard() // ШАГ
	return f.Collector.GetComputer()
}

func main() {
	// Определяем комплектации для будущих производств
	hpCollector := GetCollector(HpCollectorType)
	asusCollector := GetCollector(AsusCollectorType)

	// Создаём завод. Базово он создаёт HP
	factory := NewFactory(hpCollector)
	currentComputer := factory.CreateComputer()
	currentComputer.Print() // Выводим информацию о компьютере

	// Теперь мы хотим чтобы завод производил АСУС
	factory.SetCollecor(asusCollector)
	currentComputer = factory.CreateComputer()
	currentComputer.Print() // Выводим информацию о компьютере

	// Теперь снова заказ для завода от HP
	factory.SetCollecor(hpCollector)
	currentComputer = factory.CreateComputer()
	currentComputer.Print() // Выводим информацию о компьютере
}
