package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string
	var quit string

	for {
		fmt.Print("Enter first number: ")
		fmt.Scanln(&number1)
		_, err := fmt.Scanln(&number2)
		if err != nil {
			fmt.Errorf("invalid input number1")
		}
		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&number2)
		if err != nil {
			fmt.Errorf("invalid input number2")
		}

		fmt.Print("Enter operator (+ - * /): ")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Printf("%f %s %f = %f\n", number1, operator, number2, number1+number2)
		case "-":
			fmt.Printf("%f %s %f = %f\n", number1, operator, number2, number1-number2)
		case "*":
			fmt.Printf("%f %s %f = %f\n", number1, operator, number2, number1*number2)
		case "/":
			if number2 == 0.0 {
				fmt.Println("Error: Division by zero")
			} else {
				fmt.Printf("%f %s %f = %f\n", number1, operator, number2, number1/number2)
			}
		default:
			fmt.Println("Invalid operator")
		}

		print("Exit press q:  ")
		fmt.Scanln(&quit)

		if quit == ("q") {
			print("===============================")
			break
		}
	}
}
