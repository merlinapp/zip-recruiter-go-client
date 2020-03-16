package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	parserError        = "could not parse error"
	unauthorizedAccess = "unauthorized access"
	baseUrl            = "https://api.ziprecruiter.com/jobs/v1"
)

type ZipClient struct {
	BaseUrl string
	ApiKey  string
}

func NewZipClient() *ZipClient {
	apiKey := os.Getenv("ZIP_RECRUITER_KEY")
	if len(apiKey) == 0 {
		panic("ZIP_RECRUITER_KEY env variable not set")
	}
	return &ZipClient{BaseUrl: baseUrl, ApiKey: apiKey}
}

func (z *ZipClient) Get(request ZipRequest) (*ZipResponse, error) {
	url := fmt.Sprintf("%s?search=%s&locatiion=%s&radius_miles=%d&days_ago=%d&jobs_per_page=%d&page=%d&refined_salary=%d&api_key=%s",
		z.BaseUrl, request.Search, request.Location, request.RadiusMiles, request.Page, request.JobsPerPage, request.DaysAgo, request.RefineSalary, z.ApiKey)
	req, errNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if errNewRequest != nil {
		return nil, errNewRequest
	}
	res, errCall := http.DefaultClient.Do(req)
	if errCall != nil {
		return nil, errCall
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(unauthorizedAccess)
	}
	var zipResponse ZipResponse
	parserErr := json.Unmarshal(body, &zipResponse)
	if parserErr != nil {
		return nil, errors.New(parserError)
	}
	return &zipResponse, nil
}
