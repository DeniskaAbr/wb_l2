package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

}

// https://github.com/kgantsov/gogrep
// https://willdemaine.ghost.io/grep-from-first-principles-in-golang/
// https://pkg.go.dev/github.com/u-root/u-root/cmds/core/grep
// https://healeycodes.com/beating-grep-with-go
// https://vorozhko.net/very-simple-grep-tool-in-go-search-substring-in-files
