/*2.	Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
●	"a4bc2d5e" => "aaaabccddddde"
●	"abcd" => "abcd"
●	"45" => "" (некорректная строка)
●	"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
●	qwe\4\5 => qwe45 (*)
●	qwe\45 => qwe44444 (*)
●	qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// по правильному бы использовать лексический анализ как в примере по адресу https://github.com/bbuck/go-lexer
// но я пока затрудняюсь это сделать (хотя это и проще)
func main() {

	fmt.Println(StringUnpacker("a4bc2d5e"))
	fmt.Println(StringUnpacker("abcd"))
	fmt.Println(StringUnpacker("45"))
	fmt.Println(StringUnpacker(""))

	fmt.Println(StringUnpacker(`qwe\4\5`)) // qwe45
	fmt.Println(StringUnpacker(`qwe\45`))  // qwe44444
	fmt.Println(StringUnpacker(`qwe\\5`))  // qwe\\\\\
}

func StringUnpacker(s string) (string, error) {
	var outBuilder strings.Builder
	var previousR rune
	previousR = 0
	i := 0
	for i < len(s) {
		curRune := rune(s[i])
		if unicode.IsNumber(curRune) && i == 0 {
			return "", errors.New("string not correct")
		}
		if unicode.IsLetter(curRune) {
			previousR = curRune
			outBuilder.WriteRune(curRune)
		}
		if unicode.IsNumber(curRune) && curRune != '\\' {
			if previousR != '0' {
				n, err := strconv.Atoi(string(curRune))
				if err != nil {
					return "", err
				}
				for i := 0; i < n-1; i++ {
					outBuilder.WriteRune(previousR)
				}
			}
		}
		if curRune == rune('\\') {
			sn, n, err := Escapes(&s, i)
			if err != nil {
				return "", err
			}
			outBuilder.WriteString(sn)
			i = n
		}
		i++
	}
	return outBuilder.String(), nil
}

func Escapes(s *string, i int) (string, int, error) {
	var outBuilder strings.Builder
	if []rune(*s)[i+1] == '\\' {
		if unicode.IsNumber([]rune(*s)[i+2]) {
			n, err := strconv.Atoi(string([]rune(*s)[i+2]))
			if err != nil {
				return "", 0, err
			}
			for i := 0; i < n; i++ {
				outBuilder.WriteRune('\\')
			}
			return outBuilder.String(), i + 2, nil
		}

	} else if unicode.IsNumber([]rune(*s)[i+1]) {
		if i < len(*s)-2 && unicode.IsNumber([]rune(*s)[i+2]) {
			n, err := strconv.Atoi(string([]rune(*s)[i+2]))
			if err != nil {
				return "", 0, err
			}
			for j := 0; j < n; j++ {
				outBuilder.WriteRune([]rune(*s)[i+1])
			}
			return outBuilder.String(), i + 3, nil
		}
		outBuilder.WriteRune([]rune(*s)[i+1])
		return outBuilder.String(), i + 1, nil
	}
	return "", 0, errors.New("wrong string")
}
