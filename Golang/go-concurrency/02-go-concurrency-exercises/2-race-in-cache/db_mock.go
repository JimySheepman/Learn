package main

import "time"

type MockDB struct {
}

func (m *MockDB) Get(key string) (string, error) {
	d, _ := time.ParseDuration("20ms")
	time.Sleep(d)
	return "", nil
}

func GetMockDB() *MockDB {
	return &MockDB{}
}
