package generator

import "math/rand"

type MockData struct {
	Number int
}

type MockGenerator struct{}

func NewMockGenerator() *MockGenerator {
	return &MockGenerator{}
}

func (generator *MockGenerator) GenerateNumber() MockData {
	return MockData{Number: rand.Intn(100)}
}
