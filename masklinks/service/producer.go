package service

import (
	"os"
	"strings"
)
	type FileProducer struct {
		path string
	}
	func NewFileProducer(path string) Producer {
		return FileProducer {path:path}
	}
	func(f FileProducer) Produce() ([]string,error) {
		file,err := os.ReadFile(f.path)
		if err != nil {
			return nil, err
		}  
		return strings.Split(string(file) , "\n"),nil
		
	}