// recognize
package main

import (
	"fmt"
)

func recogize() {
	fmt.Println("recognize")
}

func answer(q string) string {
	answertext := "Can't get you! Please input: 'average age' / 'average weight' / 'average height' / 'percent male' / 'percent female'"
	patients := parseCsvPatients()
	switch {
	case q == "Hi" || q == "Hi!" || q == "hi" || q == "Hello" || q == "hello":
		answertext = "Hello! Please input: 'average age' / 'average weight' / 'average height' / 'percent male' / 'percent female'"
		break
	case q == "average age":
		answertext = fmt.Sprintf("%f", average(patients, 1)) + " patient average age"
		break
	case q == "average weight":
		answertext = fmt.Sprintf("%f", average(patients, 2)) + " patient average weight"
		break
	case q == "average height":
		answertext = fmt.Sprintf("%f", average(patients, 3)) + " patient average height"
		break
	case q == "percent male":
		answertext = fmt.Sprintf("%f", percent(patients, "male", 4)) + " % male"
		break
	case q == "percent female":
		answertext = fmt.Sprintf("%f", percent(patients, "female", 4)) + " % female"
		break
	}
	return answertext
}
