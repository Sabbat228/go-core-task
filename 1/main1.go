package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var intDecimal int
	var intOctal int
	var intHexadecimal int
	var floatNum float64
	var str string
	var boolean bool
	var complexNum complex64

	print(intDecimal, intOctal, intHexadecimal, floatNum, str, boolean, complexNum)

	combinedStr := CombineVariables(intDecimal, intOctal, intHexadecimal, floatNum, str, boolean, complexNum)
	fmt.Println("Объединенная строка:", combinedStr)

	runes := []rune(combinedStr)
	fmt.Println(runes)

	salt := "go-2024"
	hashHex := HashWithSalt(combinedStr, salt)
	fmt.Println(hashHex)
}

func print(intDecimal int, intOctal int, intHexadecimal int, floatNum float64, str string, boolean bool, complexNum complex64) {
	fmt.Printf("%T\n", intDecimal)
	fmt.Printf("%T\n", intOctal)
	fmt.Printf("%T\n", intHexadecimal)
	fmt.Printf("%T\n", floatNum)
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", boolean)
	fmt.Printf("%T\n", complexNum)
}

func CombineVariables(intDecimal, intOctal, intHexadecimal int, floatNum float64, str string, boolean bool, complexNum complex64) string {
	return fmt.Sprintf("%d %d %d %f %s %t %v", intDecimal, intOctal, intHexadecimal, floatNum, str, boolean, complexNum)
}

func HashWithSalt(str string, salt string) string {
	saltedInput := str[:len(str)/2] + salt + str[len(str)/2:]
	hash := sha256.Sum256([]byte(saltedInput))
	return hex.EncodeToString(hash[:])
}
