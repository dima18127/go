package main

import "fmt"

func maskLinks(message string) string {
	buf := []byte(message)
	prefix := []byte("http://")

	for i := 0; i <= len(buf)-len(prefix); i++ {

		// Проверяем "http://"
		match := true

		for k := range prefix {
			if buf[i+k] != prefix[k] {
				match = false
				break
			}
		}

		if match {
			j := i + len(prefix)

			// Маскируем ссылку
			for j < len(buf) &&
				buf[j] != ' ' &&
				buf[j] != '\n' &&
				buf[j] != '\t' {

				buf[j] = '*'
				j++
			}

			i = j - 1
		}
	}

	return string(buf)
}

func main() {
	fmt.Println(maskLinks(
		" http://com  Hello, its my page: http://com See you http://com ",
	))
}