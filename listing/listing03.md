Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:

```
Вывод программы:
<nil>
false

Интерфейс хранит в себе и тип интерфейса и тип самого значения.
Значение любого интерфейса, не только error, является nil в случае когда и значение и тип являются nil.
Функция Foo возвращает nil типа *os.PathError, результат мы сравниваем с nil типа nil, откуда и следует их неравенство.


https://habr.com/ru/company/vk/blog/463063/
https://habr.com/ru/post/449714/



```
