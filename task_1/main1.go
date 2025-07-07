package task_1

import (
	"crypto/sha256"
	"fmt"
)

func getArrayVars(numDecimal, numOctal, numHexadecimal int, float float64, str string, bool bool, complexNum complex64) []interface{} {
	return []interface{}{numDecimal, numOctal, numHexadecimal, float, str, bool, complexNum}
}

func defineVarType(arr []interface{}) {
	fmt.Println("Types of variables:")
	for _, v := range arr {
		fmt.Printf("%v: %T\n", v, v)
	}
}

func transformToStr(arr []interface{}) string {
	output := ""
	for _, v := range arr {
		output += fmt.Sprintf("%v", v)
	}
	return output
}

func transformToRuneSlice(input string) []rune {
	return []rune(input)
}

func hashWithSalt(input []rune) []byte {
	salt := []rune("go-2024")
	mid := len(input) / 2

	var output []rune
	output = append(output, input[:mid]...)
	output = append(output, salt...)
	output = append(output, input[mid:]...)

	byteData := make([]byte, len(output))
	for i := 0; i < len(byteData); i++ {
		byteData[i] = byte(output[i])
	}

	h := sha256.New()
	h.Write(byteData)
	return h.Sum(nil)
}

func Main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатеричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	arr := getArrayVars(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	defineVarType(arr)

	outputStr := transformToStr(arr)
	fmt.Println("\nConcatenated string:")
	fmt.Println(outputStr)

	outputRune := transformToRuneSlice(outputStr)
	hashedRune := hashWithSalt(outputRune)
	fmt.Println("\nHashed value with salt:")
	fmt.Printf("%x\n", hashedRune)
}
