package service

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}
type Service struct {
	prod Producer
	pres Presenter
}

func (s Service) maskLinks(message string) string {
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
func (s Service) Run() error {
	links, err := s.prod.Produce()
	if err != nil {
		return err
	}
	for i, link := range links {
		links[i] = s.maskLinks(link)
	}
	return s.pres.Present(links)
}
func NewService(prod Producer, pres Presenter) Service {
	return Service{
		prod: prod,
		pres: pres,
	}
}
