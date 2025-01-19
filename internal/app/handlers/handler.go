package handlers

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"
)

const lenShort = 5

type Handler struct {
	urlToShort map[string]string
	shortToURL map[string]string
}

func New() *Handler {
	h := &Handler{
		urlToShort: make(map[string]string),
		shortToURL: make(map[string]string),
	}

	return h
}

func (h *Handler) Shorten(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	urlBody, err := url.Parse(string(body))
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	short, ok := h.urlToShort[urlBody.String()]

	if !ok {
		for {
			short = randomString(lenShort)
			if _, ok := h.shortToURL[short]; !ok {
				break
			}
		}
	}

	h.urlToShort[urlBody.String()] = short
	h.shortToURL[short] = urlBody.String()

	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write([]byte(short)); err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) Resolve(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	u, ok := h.shortToURL[req.PathValue("id")]
	if !ok {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	res.Header().Set("Location", u)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func randomString(n int) string {
	c := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY"
	result := make([]byte, 0, n)
	for i := 0; i <= n; i++ {
		result = append(result, c[rand.Intn(len(c))])
	}
	return string(result)
}
