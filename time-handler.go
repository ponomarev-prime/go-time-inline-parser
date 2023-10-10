package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Создаем сканер для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Проходим по каждой строке ввода
	for scanner.Scan() {
		line := scanner.Text()
		convertedLine := convertDateFormats(line)
		fmt.Println(convertedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения стандартного ввода:", err)
	}
}

func convertDateFormats(input string) string {
	// Создаем регулярное выражение для поиска дат в формате Unix Nanoseconds
	regex := regexp.MustCompile(`\d{19}`)

	// Ищем все соответствия регулярному выражению
	matches := regex.FindAllString(input, -1)

	// Проходим по найденным датам и конвертируем их
	for _, match := range matches {
		// Преобразуем найденную дату в int64
		nanoTime, err := time.ParseDuration(match + "ns")
		if err == nil {
			// Преобразуем во временную метку Unix (секунды и наносекунды)
			sec := int64(nanoTime.Seconds())
			nsec := int64(nanoTime.Nanoseconds()) % int64(time.Second)

			// Создаем объект time.Time с помощью временной метки Unix
			t := time.Unix(sec, nsec)

			// Форматируем в заданный формат
			formattedDate := t.Format("2006.01.02_15.04.05.000000")

			// Заменяем найденную дату на новый формат в строке
			input = strings.Replace(input, match, formattedDate, -1)
		}
	}

	return input
}
