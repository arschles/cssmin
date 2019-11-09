package minsrv

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func Middleware(mapping Mapping, next http.Handler) http.Handler {
	m := map[string]string{}
	f := func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		rawFileName := mapping.FileNameForPath(r.URL.Path)
		hashedFileName := mapping.HashedFileName(rawFileName)
		f, err := pkger.Open(hashedFileName)
		if err != nil {
			log.Printf("OOPS! %s", err)
			return
		}
		next.ServeHTTP(recorder, r)
		w.WriteHeader(recorder.Code)
		W.Header().
		io.Copy(w, recorder)
	}

	return http.HandlerFunc(f)
}
