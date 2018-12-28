package xrid_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/reedobrien/xrid"
)

func TestXIDHandlerCreatesID(t *testing.T) {
	var content = []byte("foo")
	dh := DummyHandler{content}
	h := xrid.Handler(dh)
	rw := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "http://nope", nil)
	ok(t, err)
	h.ServeHTTP(rw, r)
	equals(t, rw.Body.Bytes(), content)
	equals(t, len(rw.Header().Get("X-Request-ID")), 20)
}

func TestXIDHandlerPreservesID(t *testing.T) {
	var (
		content = []byte("foo")
		header  = "myheader"
	)

	dh := DummyHandler{content}
	h := xrid.Handler(dh)
	rw := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "http://nope", nil)
	r.Header.Set("X-Request-ID", header)
	ok(t, err)
	h.ServeHTTP(rw, r)
	equals(t, rw.Body.Bytes(), content)
	equals(t, rw.Header().Get("X-Request-ID"), header)
}

type DummyHandler struct {
	content []byte
}

func (d DummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(d.content))
}

// equals fails the test if got is not equal to want.
func equals(tb testing.TB, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\n\n\twant: %#v\033[39m\n\n", filepath.Base(file), line, got, want)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}
