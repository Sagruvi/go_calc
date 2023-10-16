package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	operands := parseString(input)
	operators := parseOp(input)

	if len(operands) == 0 || len(operators) > 4 {
		fmt.Println("Ошибка: неверное количество операндов или операторов.")
		return
	}

	operand1, operand2 := operands[0], operands[1]
	operator := operators[0]

	isArabic := isArabicNumber(operand1) && isArabicNumber(operand2)
	isRoman := isRomanNumber(operand1) && isRomanNumber(operand2)

	if !isArabic && !isRoman {
		fmt.Println("Ошибка: неверный формат чисел.")
		return
	}

	var num1, num2 int
	if isArabic {
		num1, _ = strconv.Atoi(operand1)
		num2, _ = strconv.Atoi(operand2)
	} else if isRoman {
		num1 = toArabic(operand1)
		num2 = toArabic(operand2)
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
	default:
		fmt.Println("Ошибка: неверный оператор.")
		return
	}

	if isArabic {
		if result < 0 {
			panic("Результат отрицательный")
		}
		fmt.Printf("%d\n", result)
	} else if isRoman {
		romanResult := toRoman(result)
		fmt.Printf("%s\n", romanResult)
	}
}

func isArabicNumber(input string) bool {
	num, err := strconv.Atoi(input)
	return err == nil && num >= 1 && num <= 10
}

func isRomanNumber(input string) bool {
	_, exists := romanNumerals[input]
	return exists
}

func toArabic(roman string) int {
	return romanNumerals[roman]
}

func toRoman(arabic int) string {
	if arabic <= 0 {
		return "Ошибка: результат меньше 1."
	}

	romanNumerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	romanNumeral := ""

	for i, value := range arabicValues {
		for arabic >= value {
			arabic -= value
			romanNumeral += romanNumerals[i]
		}
	}

	return romanNumeral
}
func parseOp(expression string) []string {
	operators := "1234567890VIX"
	substrings := strings.FieldsFunc(expression, func(r rune) bool {
		return strings.ContainsRune(operators, r)
	})
	for i := range substrings {
		substrings[i] = strings.TrimSpace(substrings[i])
	}
	if len(substrings) > 2 {
		panic("Неверное количество операторов")
	}
	return substrings
}
func parseString(expression string) []string {
	operators := "+-*/"
	newexpression := strings.Replace(expression, " ", "", -1)
	substrings := strings.FieldsFunc(newexpression, func(r rune) bool {
		return strings.ContainsRune(operators, r)
	})
	for i := range substrings {
		substrings[i] = strings.TrimSpace(substrings[i])
	}

	return substrings
}
