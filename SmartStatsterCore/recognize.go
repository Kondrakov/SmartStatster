// recognize
package main

import (
	"fmt"
)

func recogize() {
	fmt.Println("recognize")
}

func answer(q string) string {
	answertext := "Can't get you! Please input: 'average age' / 'average height' / 'percent male' / 'percent female'"
	switch {
	case q == "Hi" || q == "Hi!" || q == "hi" || q == "Hello" || q == "hello":
		answertext = "Hello! Please input: 'average age' / 'average height' / 'percent male' / 'percent female'"
		break
	case q == "average age":
		answertext = "age"
		break
	case q == "average height":
		answertext = "height"
		break
	case q == "percent male":
		answertext = "male"
		break
	case q == "percent female":
		answertext = "female"
		break
	}
	return answertext
}
