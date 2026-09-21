package sensor

func Status(celsius, threshold float64) string {
	if celsius >= threshold {
		return "WARNING"
	}
	return "OK"
}
