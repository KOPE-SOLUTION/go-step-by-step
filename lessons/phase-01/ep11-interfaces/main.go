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

func main() {
	show("sensor-01", FixedSensor{Celsius: 27.5})
	show("sensor-02", FailedSensor{})
	show("sensor-03", FixedSensor{Celsius: 30})
}
