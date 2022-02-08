package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_echoHandler(t *testing.T) {
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
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=test", strings.NewReader("")),
			},
			want: "test",
		},
		{
			name: "post",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader("body=test")),
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.args.r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Act
			echoHandler(tt.args.w, tt.args.r)

			// Assert
			if tt.args.w.Code != http.StatusOK {
				t.Errorf("echo handler should always return OK, but %v was returned", tt.args.w.Code)
			}

			body := tt.args.w.Body.String()
			if body != tt.want {
				t.Errorf("want %v, but %v", tt.want, body)
			}
		})
	}
}
