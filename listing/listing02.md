Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод программы:
2
1

Программа показывает работу оператора `defer` который откладывает выполнение функции до возврата из окружающей его функции. Аргументы отложенного вызова оцениваются немедленно, но вызов функции не происходит до тех пор пока не произойдет возврат из окружающей функции.
Так в примере сначала происходит вызов функции `test()` так как в этой функции `defer` откладывает выполнение анонимной функции в которой значение переменной `x` оценивается инкрементированным значением 2
Потом происходит выполнение функции `anotherTest()` в области видимости которой инициализируется переменная `x` значением 0, потом оператором `defer` в анонимной функции значение `х` оценивается как x=0+1 но это не оказывает на значение `х` из за ограниченной области видимости переменной `x` для анонимной функции.
Порядок вызовов оператора `defer` происходит в порядке от последнего к первому 

Выражение defer добавляет вызов функции после ключевого слова defer в стеке приложения. Все вызовы в стеке вызываются при возврате функции, в которой они добавлены. Поскольку вызовы помещаются в стек, они производятся в порядке от последнего к первому.
https://www.digitalocean.com/community/tutorials/understanding-defer-in-go-ru

```