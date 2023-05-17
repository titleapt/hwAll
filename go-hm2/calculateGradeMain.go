package main

func grade(score int) string {

	if score >= 80 && score <= 100 {
		return "A"
	} else if score >= 70 && score <= 79 {
		return "B"
	} else if score >= 60 && score <= 69 {
		return "C"
	} else if score >= 50 && score <= 69 {
		return "D"
	} else if score <= 49 {
		return "F"
	} else {
		return "out of range"
	}
}
