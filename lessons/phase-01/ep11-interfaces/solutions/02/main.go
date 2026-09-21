package main

import "fmt"

type Reader interface {
	Read() (float64, error)
}

type FixedSensor struct {
	Celsius float64
}

func (s FixedSensor) Read() (float64, error) {
	return s.Celsius, nil
}

type FailedSensor struct{}

func (s FailedSensor) Read() (float64, error) {
	return 0, fmt.Errorf("simulated read failure")
}

func show(name string, reader Reader) {
	value, err := reader.Read()
	if err != nil {
		fmt.Printf("%s: ERROR: %v\n", name, err)
		return
	}
	fmt.Printf("%s: %.1f C\n", name, value)
}

type OffsetSensor struct {
	Base   float64
	Offset float64
}

func (s OffsetSensor) Read() (float64, error) {
	return s.Base + s.Offset, nil
}

func main() {
	show("sensor-01", FixedSensor{Celsius: 27.5})
	show("sensor-02", FailedSensor{})
	show("sensor-03", OffsetSensor{Base: 30, Offset: -1.5})
}
