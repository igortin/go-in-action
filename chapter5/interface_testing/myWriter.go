package myWriter

type MyWriter struct{}

func (w *MyWriter) Write(b []byte) error {
	return nil
}
