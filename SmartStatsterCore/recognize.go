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
	switch q {
	case "Hi", "Hi!", "hi", "Hello", "hello":
		answertext = "Hello! Please input: 'average age' / 'average weight' / 'average height' / 'percent male' / 'percent female'"
		break
	case "average age":
		answertext = fmt.Sprintf("%f", average(patients, 1)) + " patient average age"
		break
	case "average weight":
		answertext = fmt.Sprintf("%f", average(patients, 2)) + " patient average weight"
		break
	case "average height":
		answertext = fmt.Sprintf("%f", average(patients, 3)) + " patient average height"
		break
	case "percent male":
		answertext = fmt.Sprintf("%f", percent(patients, "male", 4)) + " % male"
		break
	case "percent female":
		answertext = fmt.Sprintf("%f", percent(patients, "female", 4)) + " % female"
		break
	}
	return answertext
}
