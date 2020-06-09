package tests

import (
	"github.com/louisevanderlith/theme/handles"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileDownload_ContentType_Set(t *testing.T) {
	req, err := http.NewRequest("GET", "/asset/css/bundle.css", nil)

	if err != nil {
		t.Fatal(err)
	}

	ttxt := make(chan string)

	handle := handles.SetupRoutes("localhost", "secret", "http://localhost:8086")
	for i := 0; i < 50; i++ {
		go func() {
			rr := httptest.NewRecorder()
			handle.ServeHTTP(rr, req)

			if len(rr.Header().Get("Content-Type")) == 0 {
				ttxt <- "Content-Type not Found"
				t.Fatal("Content-Type not Found")
			}
			t.Log(rr.Header().Get("Content-Type"))
			t.Fail()
			if rr.Header().Get("Content-Type") != "text/css" {
				ttxt <- rr.Body.String()
				t.Fatal(rr.Body.String())
			}
		}()
	}

	t.Log(<-ttxt)
}
