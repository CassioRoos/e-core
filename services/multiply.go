package services

type MultiplyService interface {
	GetMultiplication(records [][]string) (int64, error)
}

type multiply struct {
}

func NewMultiplyService() MultiplyService {
	return &multiply{}
}

func (m *multiply) Action(j, i int64)int64{
	total := j * i
	return total
}


func (m *multiply) GetMultiplication(records [][]string) (int64, error) {
	return processInt64(records,1, m)
}
