package modules

import (
	"net/http"
)

type ResponseObserver struct {
	http.ResponseWriter
	tracerSpanClosed bool
	MaiStatus        int
	MaiWritten       int64
	wroteHeader      bool
	maiRes           []byte
}

func (o *ResponseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	if !o.tracerSpanClosed {
		o.tracerSpanClosed = true
	}
	n, err = o.ResponseWriter.Write(p)
	o.MaiWritten += int64(n)
	o.maiRes = p
	return
}

func (o *ResponseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		if !o.tracerSpanClosed {
			o.tracerSpanClosed = true
		}
		return
	}
	o.wroteHeader = true
	o.MaiStatus = code

	if !o.tracerSpanClosed {
		o.tracerSpanClosed = true
	}

}

func (o *ResponseObserver) GetMaiRes() []byte {
	return o.maiRes
}
