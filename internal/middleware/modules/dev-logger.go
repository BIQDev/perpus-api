package modules

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func DevLoggerInit(r *http.Request, o *ResponseObserver) {
	logger := log.New(os.Stderr, "", 0)

	addr := r.RemoteAddr
	if i := strings.LastIndex(addr, ":"); i != -1 {
		addr = addr[:i]
	}

	logger.Printf("%s - -  %q %d %d",
		addr,
		fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
		o.MaiStatus,
		o.MaiWritten,
	)


	/*
		logger.Printf("%s - - [%s] %q %d %d %q %q",
			addr,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			o.status,
			o.written,
			r.Referer(),
			r.UserAgent(),
		)*/
}
