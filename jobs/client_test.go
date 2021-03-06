package jobs

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
	url := fmt.Sprintf("%s%s", srv.URL, "/")
	zipClient := NewZipClient(url, "key")
	jobs, err := zipClient.Get(ZipRequest{
		Search:       "cashier",
		Location:     "Brooklyn, NY",
		RadiusMiles:  0,
		Page:         0,
		JobsPerPage:  0,
		DaysAgo:      0,
		RefineSalary: 0,
	})
	s.NoError(err)
	s.NotNil(jobs)
}

func (s *zipClientSuite) TestZipClient_GetJobs_Failed() {
	srv := newServerForGetFailed500()
	defer srv.Close()
	url := fmt.Sprintf("%s%s", srv.URL, "/")
	zipClient := NewZipClient(url, "key")
	jobs, err := zipClient.Get(ZipRequest{})
	s.Error(err)
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
