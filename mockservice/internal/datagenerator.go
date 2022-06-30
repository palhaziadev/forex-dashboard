package service

import "math/rand"

type MockData struct {
	Number int
}

func GenerateNumber() MockData {
	return MockData{Number: rand.Intn(100)}
}
