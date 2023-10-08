package main

import (
	"fmt"
	"os"
	"trucode.com/challenge1/bmi"
	"trucode.com/challenge1/evenodd"
	"trucode.com/challenge1/foo"
	"trucode.com/challenge1/mario"
	equations "trucode.com/challenge1/ohm"
)

func main() {
	args := os.Args
	userData := args[1]

	switch userData {

	case "evenodd":
		var num int
		fmt.Println("Write the number you want to evaluate to find out if it is even or odd")
		fmt.Scan(&num)
		fmt.Println("The number you entered is:")
		fmt.Println(evenodd.Evenodd(num))

	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Write the values ​​for V, R and I")
		fmt.Scan(&v, &r, &i)
		if v == 0 && r > 0 && i > 0 {
			fmt.Println("The voltage value in volts is:")
		} else if v > 0 && r == 0 && i > 0 {
			fmt.Println("The resistance value in ohms is:")
		} else if v > 0 && r > 0 && i == 0 {
			fmt.Println("The value of the current in amperes is:")
		}else if v > 0 && r > 0 && i > 0 {
			fmt.Println("You already have all the values")
		}else{
			fmt.Println("You must enter at least two values")
		}
		fmt.Println(equations.Ohm(v, r, i))

	case "foobar":
		var limit int
		fmt.Println("Write the limit number for which you want to evaluate Foobar")
		fmt.Scan(&limit)
		fmt.Println("Here is the result")
		fmt.Println(foo.Foobar(limit))
	
		case "bmi":
		var h float32
		var w float32
		fmt.Println("How much do you weigh in kilograms? (don't lie)")
		fmt.Scan(&w)
		fmt.Println("How tall are you in meters? (barefoot)")
		fmt.Scan(&h)
		fmt.Println("Your bmi is:")
		fmt.Println(bmi.Bmi(w, h))
		index := bmi.Bmi(w, h);
		if  index < 18.5 {
			fmt.Println("You are underweight, add more potato to the broth")
		} else if index >= 18.5 && index < 25 {
			fmt.Println("You have a normal weight, I have healthy envy of you'")
		} else if index >= 25 {
			fmt.Println("You are overweight, I know, the pandemic has affected us all")
		}

	case"mario":
		for{
			var height int
			fmt.Println("Write the height number of the pyramid you want to build:")
			fmt.Scan(&height)
			
			if height>=1 && height<=8 {
				fmt.Println("Here is the result")
				fmt.Println(mario.Mario(height))
			break
			}else {
				fmt.Println("Please enter a number between 1 and 8")
			}

		}
		
	}
}
