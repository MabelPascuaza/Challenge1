package mario

func Mario(height int) string {
var pyramid string
	for i := 1; i <= height; i++ {		
		for j := 1; j <= height-i; j++ {
			pyramid += " "
		}
		for j := 1; j <= i; j++ {
			pyramid += "#"
		}
		pyramid += "\n"
	} 
	return pyramid
}
