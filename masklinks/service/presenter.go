package service
import (
	"os"
	"strings"
)
type SendFile struct {
		path string
	}
func NewSendFile (path string) Presenter {
	return SendFile{path: path}
}
func (f SendFile) Present (text [] string)  error {
		data := strings.Join(text, "\n")
		return os.WriteFile(f.path, []byte(data), 0644)
	}