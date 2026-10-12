package main

import (
	"fmt"
	"masklinks/service"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажи путь к исходному файлу")
		return
	}
	InputPath := os.Args[1]
	OutputPath := "MaskedData"
	if len(os.Args) >= 3 {
		OutputPath = os.Args[2]
	}
	service := service.NewService(
		service.NewFileProducer(InputPath),
		service.NewSendFile(OutputPath),
	)
	if err := service.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
