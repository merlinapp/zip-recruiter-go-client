package zip_recruiter_go_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.ziprecruiter.com/jobs/v1"
)

type ZipClient struct {
	BaseUrl string
}

func NewZipClient() *ZipClient {
	return &ZipClient{BaseUrl: baseUrl}
}

func (z *ZipClient) Get(apiKey string,
	search string,
	location string,
	radiusMiles string,
	page string,
	jobsPerPage string,
	daysAgo string,
	refineSalary string,
) ([]byte, int, error) {

	url := fmt.Sprintf("%s?search=%s&locatiion=%s&radius_miles=%s&days_ago=%s&jobs_per_page=%s&page=%s&refined_salary=%s&api_key=%s",
		z.BaseUrl, search, location, radiusMiles, page, jobsPerPage, daysAgo, refineSalary, apiKey)
	req, errNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if errNewRequest != nil {
		return nil, http.StatusInternalServerError, errNewRequest
	}
	res, errCall := http.DefaultClient.Do(req)
	if errCall != nil {
		return nil, http.StatusInternalServerError, errCall
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	if res.StatusCode != 200 {
		return nil, res.StatusCode, nil
	}
	return body, res.StatusCode, nil
}
