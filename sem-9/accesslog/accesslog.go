//go:build !solution

package accesslog

import (
	"context"
	"crypto/rand"
	"net/http"
)

type Entry struct {
	ID     string
	Method string
	Path   string
	Status int
	Bytes  int
}

type ctxKey struct{}

type recorder struct {
	http.ResponseWriter
	status int
	bytes  int
	wrote  bool
}

func (w *recorder) WriteHeader(code int) {
	if !w.wrote {
		w.status = code
		w.wrote = true
		w.ResponseWriter.WriteHeader(code)
	}
}

func (w *recorder) Write(p []byte) (int, error) {
	w.wrote = true
	n, err := w.ResponseWriter.Write(p)
	w.bytes += n
	return n, err
}

func Wrap(next http.Handler, log func(Entry)) http.Handler {
	// TODO: attach the ID and record the response.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = rand.Text()
		}
		w.Header().Set("X-Request-ID", id)

		rec := &recorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r.WithContext(context.WithValue(r.Context(), ctxKey{}, id)))

		log(Entry{
			ID:     id,
			Method: r.Method,
			Path:   r.URL.Path,
			Status: rec.status,
			Bytes:  rec.bytes,
		})
	})
}

func RequestID(ctx context.Context) string {
	id, _ := ctx.Value(ctxKey{}).(string)
	return id
}
