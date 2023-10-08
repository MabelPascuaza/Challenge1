package equations

func Ohm(v, r, i float32) float32 {

	if v == 0 && r > 0 && i > 0 {
		// Calcular voltaje 
		return i * r
	} else if v > 0 && r == 0 && i > 0 {
		// Calcular resistencia 
		return v / i
	} else if v > 0 && r > 0 && i == 0 {
		// Calcular corriente 
		return v / r
	} else if v > 0 && r > 0 && i > 0 {
		return 0
	}
	return 0
}
