package jobs

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	response, _ = ioutil.ReadFile("fixtures/response.json")
)

type zipClientSuite struct {
	suite.Suite
}

func (s *zipClientSuite) SetupTest() {
	_ = os.Setenv("ZIP_RECRUITER_KEY", "A")
}
func TestService(t *testing.T) {
	suite.Run(t, &zipClientSuite{})
}

func (s *zipClientSuite) TestZipClient_GetJobs_Succeed() {
	srv := newServerForGetSuccess()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, err := zipClient.Get(ZipRequest{})
	s.NoError(err)
	s.NotNil(jobs)
}
func (s *zipClientSuite) TestZipClient_GetJobs_Unauthorized() {
	srv := newServerForGetUnauthorized()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, err := zipClient.Get(ZipRequest{})
	s.NoError(err)
	s.Nil(jobs)
}
func (s *zipClientSuite) TestZipClient_GetJobs_Failed() {
	srv := newServerForGetFailed500()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, err := zipClient.Get(ZipRequest{})
	s.NoError(err)
	s.Nil(jobs)
}

func newServerForGetSuccess() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("some error"))
		}
		w.WriteHeader(200)
		w.Write(response)
	}))
}

func newServerForGetFailed500() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Server Panic"))
	}))
}

func newServerForGetUnauthorized() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		w.Write([]byte("{\"error\":\"unauthorized\"}"))
	}))
}
