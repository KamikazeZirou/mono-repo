package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8080", strings.NewReader("")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloHandler(tt.args.w, tt.args.r)
			if tt.args.w.Code != http.StatusOK {
				t.Errorf("echo handler should always return OK, but %v was returned", tt.args.w.Code)
			}
		})
	}
}
