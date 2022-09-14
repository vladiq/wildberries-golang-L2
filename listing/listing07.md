Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет содержимое каналов a и b, а потом бесконечно будет выводить 0.
Так как канал c не закрыт, то чтение из него будет возвращать нулевое значение типа int (0).
Для корректной работы программы можно внутри select в merge() добавить проверку на то, не закрыты ли каналы a и b.
Есть удобная проверка вида case v, ok := <-a
Если каналы a и b закрыты, то можно закрыть канал c и завершить выполнение пишущей в канал c горутины, что успешно остановит цикл for в main().
```