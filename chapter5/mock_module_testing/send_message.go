package msg

type Message struct {
	//
}

func (m *Message) Send(email string, subject string, body []byte) error {
	//
	return nil
}

type Messanger interface {
	Send(email string, subject string, body []byte) error
}

func Alert(m Messanger, email string, subject string, problem []byte) error {
	return m.Send(email, subject, problem)
}
