package sensor

func Status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}

func ToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}
