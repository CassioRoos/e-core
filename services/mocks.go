package services

import "github.com/stretchr/testify/mock"

type EchoServiceMock struct {
	mock.Mock
}

type SumServiceMock struct {
	mock.Mock
}

type MultiplyServiceMock struct {
	mock.Mock
}

type FlattenServiceMock struct {
	mock.Mock
}

type InvertServiceMock struct {
	mock.Mock
}

func (e *EchoServiceMock)GetEcho(records [][]string) string{
	args := e.Called(records)
	return args.String(0)
}

func (s *SumServiceMock) GetSum(records [][]string) (int64,error){
	args := s.Called(records)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MultiplyServiceMock) GetMultiplication(records [][]string) (int64, error){
	args := m.Called(records)
	return args.Get(0).(int64), args.Error(1)
}

func (f *FlattenServiceMock) GetFlatten(records [][]string) string{
	args := f.Called(records)
	return args.String(0)
}

func (i *InvertServiceMock)GetInvert(records [][]string) string{
	args := i.Called(records)
	return args.String(0)
}