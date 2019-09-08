package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/theme/routers"
)

var (
	epox resins.Epoxi
)

func init() {
	srvc := bodies.NewService("Artifact.API", "/certs/none.pem", 8093, servicetype.API)
	srvc.ID = "Tester1"

	epox = resins.NewMonoEpoxy(srvc)
	routers.Setup(epox)
	epox.EnableCORS(".localhost/")
}

func TestFileDownload_ContentType_Set(t *testing.T) {
	req, err := http.NewRequest("GET", "/asset/css/bundle.css", nil)

	if err != nil {
		t.Fatal(err)
	}

	ttxt := make(chan string)

	handle := epox.GetRouter()
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
