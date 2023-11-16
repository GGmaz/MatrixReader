package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrixKat, ok := createMatrixFromInput("fsd.txt")
	if !ok {
		return
	}

	matrixInt, ok := createMatrixFromInput("mesnatost.txt")
	if !ok {
		return
	}

	x := 21
	y := 63

	valueKategorija := getElement(matrixKat, x, y)
	valueMesnatost := getElement(matrixInt, x, y)

	fmt.Printf("Kategorija Matrix[%d][%d] = %s\n", x, y, valueKategorija)
	fmt.Printf("Mesnatost Matrix[%d][%d] = %s\n", x, y, valueMesnatost)
}

func createMatrixFromInput(filename string) ([][]string, bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, false
	}
	defer file.Close()

	var matrixStr [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		var row []string
		for _, v := range values {
			num := v
			row = append(row, num)
		}

		matrixStr = append(matrixStr, row)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, false
	}
	return matrixStr, true
}

func getElement(matrix [][]string, x, y int) string {
	// Check if the provided indices are within the valid range
	if x < 5 || x > 35 || y < 50 || y > 89 {
		fmt.Println("Indices out of range")
		return "O"
	}

	return matrix[x-5][y-50]
}

func printMatrixStr(matrix [][]string) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	}
}

func replaceZeros(matrix [][]string) [][]string {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i < 2 || (i == 2 && j > 2) || (i == 3 && j > 6) || (i == 4 && j > 11) || (i == 5 && j > 15) || (i == 6 && j > 20) || (i == 7 && j > 25) || (i == 8 && j > 30) || (i == 9 && j > 35) {
				matrix[i][j] = "S"
			} else if i < 12 {
				matrix[i][j] = "E"
			} else if (i == 12 && j > 0) || (i == 13 && j > 4) || (i == 14 && j > 8) || (i == 15 && j > 12) || (i == 16 && j > 16) || (i == 17 && j > 20) || (i == 18 && j > 24) || (i == 19 && j > 28) || (i == 20 && j > 32) || (i == 21 && j > 36) {
				matrix[i][j] = "E"
			} else if i < 23 {
				matrix[i][j] = "U"
			} else if (i == 23 && j > 1) || (i == 24 && j > 5) || (i == 25 && j > 9) || (i == 26 && j > 13) || (i == 27 && j > 17) || (i == 28 && j > 21) || (i == 29 && j > 25) || (i == 30 && j > 29) || (i == 31 && j > 33) || (i == 32 && j > 37) {
				matrix[i][j] = "U"
			} else {
				matrix[i][j] = "R"
			}
		}
	}
	return matrix
}
