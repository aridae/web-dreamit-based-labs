package testinghelp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// мы тестируем методы хэндлера, которые принимают w http.ResponseWriter, r *http.Request
// нам нужна хэлпер функция, которая создает реквест по входным параметрам теста -- как раз билдер понадобится
// и функция, котоая проверяет, что записано в респонс райтер -- чтобы сравнить с ожидаемым

// http.MethodGet
type RequestBuilder struct {
	method      string
	body        []byte
	route       string
	queryParams map[string]string
	routeVars   map[string]string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		queryParams: make(map[string]string, 0),
		routeVars:   make(map[string]string, 0),
	}
}

func (b *RequestBuilder) Build() (*http.Request, *httptest.ResponseRecorder) {
	r := httptest.NewRequest(b.method, b.route, io.NopCloser(bytes.NewBuffer(b.body)))
	r = mux.SetURLVars(r, b.routeVars)
	q := r.URL.Query()
	for k, v := range b.queryParams {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()

	return r, w
}

func (b *RequestBuilder) WithParameter(key, val string) *RequestBuilder {
	b.queryParams[key] = val
	return b
}

func (b *RequestBuilder) WithParameters(params map[string]string) *RequestBuilder {
	for k, v := range params {
		b.WithParameter(k, v)
	}
	return b
}

func (b *RequestBuilder) WithRouteVar(key, val string) *RequestBuilder {
	b.routeVars[key] = val
	return b
}

func (b *RequestBuilder) WithRouteVars(vars map[string]string) *RequestBuilder {
	for k, v := range vars {
		b.WithRouteVar(k, v)
	}
	return b
}

func (b *RequestBuilder) WithBody(body interface{}) *RequestBuilder {
	jsonBytes, _ := json.Marshal(body)
	b.body = jsonBytes
	return b
}

func (b *RequestBuilder) WithRoute(route string) *RequestBuilder {
	b.route = route
	return b
}

func (b *RequestBuilder) WithMethod(method string) *RequestBuilder {
	b.method = method
	return b
}
