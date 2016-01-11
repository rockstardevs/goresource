package mocks

type MockEntity struct {
	Id string
}

func (m *MockEntity) HasId() bool {
	if m.Id != "" {
		return true
	}
	return false
}

func (m *MockEntity) GetId() string {
	return m.Id
}
