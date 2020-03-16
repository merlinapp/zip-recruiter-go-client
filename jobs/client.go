package jobs

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.ziprecruiter.com/jobs/v1"
)

func Get(apiKey string,
	search string,
	location string,
	radiusMiles string,
	page string,
	jobsPerPage string,
	daysAgo string,
	refineSalary string,
) (interface{}, int, error) {

	url := fmt.Sprintf("%s?search=%s&locatiion=%s&radius_miles=%s&days_ago=%s&jobs_per_page=%s&page=%s&refined_salary=%s&api_key=%s",
		baseUrl, search, location, radiusMiles, page, jobsPerPage, daysAgo, refineSalary, apiKey)
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
	return body, res.StatusCode, nil
}
