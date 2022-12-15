package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUnpack(t *testing.T) {

	//●	"a4bc2d5e" => "aaaabccddddde"
	//●	"abcd" => "abcd"
	//●	"45" => "" (некорректная строка)
	//●	"" => ""
	//●	qwe\4\5 => qwe45 (*)
	//●	qwe\45 => qwe44444 (*)
	//●	qwe\\5 => qwe\\\\\ (*)

	assert := assert.New(t)

	var testsOK = []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}

	var testsWRONG = []struct {
		input    string
		expected string
	}{
		{"45", ""},
	}

	for i, test := range testsOK {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			expected, err := StringUnpacker(test.input)
			assert.NoError(err, nil)
			assert.Equal(expected, test.expected)
		})
	}
	for i, test := range testsWRONG {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			expected, err := StringUnpacker(test.input)
			assert.Error(err, errors.New("string not correct"))
			assert.Equal(expected, test.expected)
		})
	}
}
