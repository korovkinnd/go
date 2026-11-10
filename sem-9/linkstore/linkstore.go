//go:build !solution

package linkstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

const maxBody = 1024

type store struct {
	mu    sync.Mutex
	byURL map[string]string
	byID  map[string]string
	next  int
}

type createRequest struct {
	URL string `json:"url"`
}

func New() http.Handler {
	// TODO: register routes and store links.
	s := &store{byURL: map[string]string{}, byID: map[string]string{}}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /links", s.create)
	mux.HandleFunc("GET /l/{id}", s.redirect)
	return mux
}

func (s *store) create(w http.ResponseWriter, r *http.Request) {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxBody))
	var tooLarge *http.MaxBytesError
	if errors.As(err, &tooLarge) {
		http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
		return
	}
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	var req createRequest
	if err := dec.Decode(&req); err != nil || dec.Decode(&struct{}{}) != io.EOF {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if !validURL(req.URL) {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	id, ok := s.byURL[req.URL]
	status := http.StatusOK
	if !ok {
		s.next++
		id = strconv.FormatInt(int64(s.next), 36)
		s.byURL[req.URL] = id
		s.byID[id] = req.URL
		status = http.StatusCreated
	}
	s.mu.Unlock()

	w.Header().Set("Location", "/l/"+id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (s *store) redirect(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	dest, ok := s.byID[r.PathValue("id")]
	s.mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Location", dest)
	w.WriteHeader(http.StatusFound)
}

func validURL(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	return (u.Scheme == "http" || u.Scheme == "https") && u.Hostname() != "" && u.User == nil
}
