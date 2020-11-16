package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/BIQDev/perpus-api/internal/middleware/modules"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func (m *biqMiddleware) Logger(ro *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			o := &modules.ResponseObserver{
				ResponseWriter: w,
			}

			contentLength := r.Header.Get("Content-Length")
			clExceedMax := false
			if contentLength != "" {

				contentSize, errContentSize := strconv.Atoi(contentLength)
				contentSizeMax, errContentSizeMax := strconv.Atoi(os.Getenv("MAI_LOG_SIZE_MAX"))

				clExceedMax = errContentSize != nil || errContentSizeMax != nil || contentSize > contentSizeMax
			}
			var bodyBytes []byte
			if !clExceedMax && r.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(r.Body)
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			if clExceedMax {
				emptyBody := struct {
					Message string `json:"message"`
				}{
					Message: "Body size exceed the limit of: " + os.Getenv("MAI_LOG_SIZE_MAX") + " bytes",
				}
				bodyBytes, _ = json.Marshal(emptyBody)
				w.Write(bodyBytes)
				return
			}

			next.ServeHTTP(o, r)

			if os.Getenv("GO_ENV") == "development" {
				modules.DevLoggerInit(r, o)
			}
		})
	}
}
