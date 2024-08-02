package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []string{
	"", "I", "II", "III", "IV", "V",
	"VI", "VII", "VIII", "IX", "X",
}

func romanToInteger(roman string) (int, error) {
	if value, exists := romanToInt[roman]; exists {
		return value, nil
	}
	return 0, errors.New("неверная римская цифра1 ")
}

func integerToRoman(num int) (string, error) {
	if num <= 0 || num >= len(intToRoman) {
		return "", errors.New("результат выходит за пределы допустимого диапазона для римских цифр")
	}
	return intToRoman[num], nil
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("недопустимый оператор")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("недопустимый формат: выражение должно быть в форме 'a + b'")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	aArabic, errA := strconv.Atoi(aStr)
	bArabic, errB := strconv.Atoi(bStr)
	isRoman := errA != nil && errB != nil

	var a, b int
	var err error

	if isRoman {
		a, err = romanToInteger(aStr)
		if err != nil {
			panic(err)
		}
		b, err = romanToInteger(bStr)
		if err != nil {
			panic(err)
		}
	} else {
		if errA != nil || errB != nil {
			panic("смешанные системы счисления не допускаются")
		}
		a = aArabic
		b = bArabic
	}

	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		panic("числа должны быть в диапазоне от 1 до 10")
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	if isRoman {
		if result <= 0 {
			panic("результат находится вне диапазона для римских цифр")
		}
		romanResult, err := integerToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
