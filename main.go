package main

func maskLinks(message string) []byte { 
buf := []byte(message)
for i := 0; i < len(buf)-7; i++ {
 // Проверяем, начинается ли здесь "http://"
 if buf[i] == 'h' &&
  buf[i+1] == 't' &&
  buf[i+2] == 't' &&
  buf[i+3] == 'p' &&
  buf[i+4] == ':' &&
  buf[i+5] == '/' &&
  buf[i+6] == '/' {

  // Начинаем после "http://"
  j := i + 7

  // Ищем конец ссылки
  for j < len(buf) &&
   buf[j] != ' ' &&
   buf[j] != '\n' &&
   buf[j] != '\t' {

   buf[j] = '*'
   j++
  }

  // Перескакиваем сразу в конец замаскированной ссылки
  i = j - 1
 }
}

return buf
}

func main(){
maskLinks("Hello, its my page: http://localhost123.com See you")
}





