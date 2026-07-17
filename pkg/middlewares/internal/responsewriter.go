package internal

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (rw *ResponseWriter) StatusCode() int {
	return rw.statusCode
}

func (rw *ResponseWriter) Size() int {
	return rw.size
}
