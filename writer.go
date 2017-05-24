package proxy

import (
	"bytes"
	"net/http"
)

type WriterRecorder struct {
	http.ResponseWriter
	status int
	body   bytes.Buffer
}

func (w *WriterRecorder) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func (w *WriterRecorder) Write(body []byte) (n int, err error) {
	if n, err := w.body.Write(body); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(body)
}

func (w *WriterRecorder) Body() []byte {
	return w.body.Bytes()
}

func (w *WriterRecorder) Status() int {
	return w.status
}