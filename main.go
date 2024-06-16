package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	fields := strings.Fields(input)
	if len(fields) != 3 { // проверка ввода строки
		panic("Данные должны вводиться в одной строке")
	}
	a := fields[0] // первое число
	b := fields[1] // знак
	c := fields[2] // второе число
	numsysa := numsys(a)
	numsysc := numsys(c)
	if numsysa != numsysc { // проверка системы счисления
		panic("Введены числа разных систем счисления")
	}
	fmt.Print(calc(trans(a), trans(c), b, numsysa))
}

func numsys(n string) string { // проверка системы счисления
	numsys := ""
	switch n {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		numsys = "arabic"
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		numsys = "roman"
	default:
		panic("Введите целое число от 1 до 10")
	}
	return numsys
}

func trans(n string) int { // перевод чисел в Int
	t := 0
	if numsys(n) == "arabic" {
		t, _ = strconv.Atoi(n)
	} else {
		switch n {
		case "I":
			t = 1
		case "II":
			t = 2
		case "III":
			t = 3
		case "IV":
			t = 4
		case "V":
			t = 5
		case "VI":
			t = 6
		case "VII":
			t = 7
		case "VIII":
			t = 8
		case "IX":
			t = 9
		case "X":
			t = 10
		}
	}
	return t
}

func calc(a, c int, b, numsysa string) string { // вычисление

	result := 0
	switch b {
	case "/":
		result = a / c
	case "*":
		result = a * c
	case "+":
		result = a + c
	case "-":
		result = a - c
	default:
		panic("Неизвестная операция")
	}
	// fmt.Println("проверка: ",result) //проверка
	var results string
	if numsysa == "roman" && result <= 0 {
		panic("В римской системе счисления нет отрицательных чисел")
	} else if numsysa == "roman" {
		resultStr := strconv.Itoa(result)
		if len(resultStr) == 3 {
			results = "C"
		} else if len(resultStr) == 2 {
			switch resultStr[0] {
			case '1':
				results = "X"
			case '2':
				results = "XX"
			case '3':
				results = "XXX"
			case '4':
				results = "XL"
			case '5':
				results = "L"
			case '6':
				results = "LX"
			case '7':
				results = "LXX"
			case '8':
				results = "LXXX"
			case '9':
				results = "XC"
			}
			switch resultStr[1] {
			case '1':
				results += "I"
			case '2':
				results += "II"
			case '3':
				results += "III"
			case '4':
				results += "IV"
			case '5':
				results += "V"
			case '6':
				results += "VI"
			case '7':
				results += "VII"
			case '8':
				results += "VIII"
			case '9':
				results += "IX"
			}
		} else if len(resultStr) == 1 {
			switch resultStr[0] {
			case '1':
				results = "I"
			case '2':
				results = "II"
			case '3':
				results = "III"
			case '4':
				results = "IV"
			case '5':
				results = "V"
			case '6':
				results = "VI"
			case '7':
				results = "VII"
			case '8':
				results = "VIII"
			case '9':
				results = "IX"
			}
		}
	} else if numsysa == "arabic" {
		results = strconv.Itoa(result)
	}
	return results
}
