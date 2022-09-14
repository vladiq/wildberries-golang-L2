Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Здесь мы так же, как и в листинге 3 сравниваем nil с интерфейсом.
Для корректного сравнения нужно использовать if err.(*customError) != nil, либо использовать reflect.
Также можно заменить var err error на var err *customError
```
