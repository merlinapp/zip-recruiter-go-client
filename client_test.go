package zip_recruiter_go_client

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	response, _ = ioutil.ReadFile("fixtures/response.json")
)

type zipClientSuite struct {
	suite.Suite
}

func TestService(t *testing.T) {
	suite.Run(t, &zipClientSuite{})
}

func (s *zipClientSuite) TestZipClient_GetJobs_Succeed() {
	srv := newServerForGetSuccess()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, status, err := zipClient.Get("", "", "", "", "", "", "", "")
	s.NoError(err)
	s.Equal(http.StatusOK, status)
	s.NotNil(jobs)
}
func (s *zipClientSuite) TestZipClient_GetJobs_Unauthorized() {
	srv := newServerForGetUnauthorized()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, status, err := zipClient.Get("", "", "", "", "", "", "", "")
	s.NoError(err)
	s.Equal(http.StatusUnauthorized, status)
	s.Nil(jobs)
}
func (s *zipClientSuite) TestZipClient_GetJobs_Failed() {
	srv := newServerForGetFailed500()
	defer srv.Close()
	zipClient := NewZipClient()
	zipClient.BaseUrl = fmt.Sprintf("%s%s", srv.URL, "/")
	jobs, status, err := zipClient.Get("", "", "", "", "", "", "", "")
	s.NoError(err)
	s.Equal(http.StatusInternalServerError, status)
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
