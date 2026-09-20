package sensor

func IsWarning(celsius float64) bool {
	return celsius >= 30
}
