package services

type SumService interface {
	GetSum(records [][]string) (int64,error)
}

type sum struct {
}

func NewSumService() SumService {
	return &sum{}
}

func (s *sum) Action(j, i int64)int64{
	total := j + i
	return total
}


func (s *sum) GetSum(records [][]string) (int64,error) {
	return processInt64(records,0, s)
}