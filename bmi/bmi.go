package bmi

func Bmi(weight, height float32) float32 {
	var result float32
	result=weight/(height*height)
	return result
}
