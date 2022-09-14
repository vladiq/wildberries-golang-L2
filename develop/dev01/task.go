package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const host = "0.ru.pool.ntp.org"

func main() {
	ntpTime, err := GetNtpTime(host)
	if err != nil {
		logger := log.New(os.Stderr, "", 1)

		logger.Println(err)
		os.Exit(1)

	}
	fmt.Println(ntpTime)
}

func GetNtpTime(host string) (time.Time, error) {
	return ntp.Time(host)
}
