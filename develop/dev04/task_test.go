package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	input := []string{"тяпка", "листок", "Пятак", "слиток", "пятка", "Ласков", "столик", "словак", "славок", "сковал", "ангола"}

	testOutput := make(map[string][]string)
	testOutput["ласков"] = append(testOutput["ласков"], "ласков", "сковал", "славок", "словак")
	testOutput["листок"] = append(testOutput["листок"], "листок", "слиток", "столик")
	testOutput["тяпка"] = append(testOutput["тяпка"], "пятак", "пятка", "тяпка")

	output := Anagram(&input)
	fmt.Println(*output)
	for k := range *output {
		if ok := reflect.DeepEqual(testOutput[k], (*output)[k]); !ok {
			t.Error("Ошибка, результаты не соответствует ожидаемому")
		}
	}
}
