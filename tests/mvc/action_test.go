package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type SomeAction struct {
}

func (a *SomeAction) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestActionAsHandler(t *testing.T) {
	Convey("Given an arbitrary action instance", t, func() {

		action := new(SomeAction)

		Convey("Then it should be possible to use that action instance as a request handler", func() {
			req, err := http.NewRequest("GET", "/", nil)

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(action.Handler)

			Convey("And the handler should behave correctly", func() {
				handler.ServeHTTP(rr, req)

				if status := rr.Code; status != http.StatusOK {
					t.Errorf("Handler returned unexpected status code: Expected %v, got %v", http.StatusOK, status)
				}
			})
		})
	})
}
