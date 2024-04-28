package main

import "fmt"

// Интерфейс зарядного устройства с Lightning портом
type LightningCharger interface {
	ChargeWithLightning()
}

// Интерфейс зарядного устройства с USB портом
type USBCharger interface {
	ChargeWithUSB()
}

// Ноутбук с Lightning портом
type MacBook struct{}

func (m *MacBook) ChargeWithLightning() {
	fmt.Println("Зарядка через порт Lightning")
}

// Адаптер для преобразования Lightning в USB
type LightningToUSBAdapter struct {
	macBook *MacBook
}

func (l *LightningToUSBAdapter) ChargeWithUSB() {
	fmt.Println("Адаптер преобразует сигнал Lightning в USB")
	l.macBook.ChargeWithLightning()
}

func main() {
	macBook := &MacBook{}
	adapter := &LightningToUSBAdapter{macBook: macBook}

	// Использование адаптера для зарядки ноутбука через USB
	adapter.ChargeWithUSB()
}
