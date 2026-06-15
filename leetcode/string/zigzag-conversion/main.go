package main

func convert(s string, numRows int) string {

	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	currentRow := 0
	goingDown := true

	for _, value := range s {
		rows[currentRow] += string(value)

		if currentRow == 0 {
			goingDown = true
		} else if currentRow == numRows-1 {
			goingDown = false
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {
	convert("ABCDWSJDSKJ", 3)
}
