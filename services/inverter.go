package services

type InvertService interface {
	GetInvert(records [][]string) string
}

type invert struct {}

func NewInvertService() InvertService {
	return &invert{}
}

func (e *invert) GetInvert(records [][]string) string {
	var response string
	for i := 0; i < len(records); i++ {
		line := ""
		for j := 0; j < len(records); j++ {
			line += records[j][i] + ","
		}
		line = line[:len(line)-1]
		response += line + "\n"
	}
	return response
}