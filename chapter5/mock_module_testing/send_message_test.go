package msg

import (
	"testing"
)

// Create Mock structure and it is suit inbterface Massanger
type MockMessage struct {
	email   string
	subject string
	body    []byte
}

// Mock method
func (m *MockMessage) Send(email string, subject string, problem []byte) error {
	m.email = email
	m.subject = subject
	m.body = problem
	return nil
}

func TestAlert(t *testing.T) {

	var tests = []struct {
		email    string
		subject  string
		body     []byte
		expected string
	}{
		{
			"team@example.com",
			"Fake Alert",
			[]byte("Some text to send it"),
			"Fake Alert",
		},
		{
			"team@example.com",
			"Fake Alert",
			[]byte("Some text to send it"),
			"Fake Alert",
		},
	}

	for _, tt := range tests {
		mock := &MockMessage{}

		/*
		 Goal to test function named the Alert from main code
		*/

		Alert(mock, tt.email, tt.subject, tt.body)
		if mock.subject != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, mock.subject)
		}
	}
}
