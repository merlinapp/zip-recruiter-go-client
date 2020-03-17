package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urlUtils "net/url"
	"os"
)

const (
	parserError        = "could not parse error"
	unauthorizedAccess = "unauthorized access"
)

type ZipClient struct {
	BaseUrl string
	ApiKey  string
}

func NewZipClient() Client {
	apiKey := os.Getenv("ZIP_RECRUITER_KEY")
	if len(apiKey) == 0 {
		panic("ZIP_RECRUITER_KEY env variable not set")
	}
	baseUrl := os.Getenv("ZIP_RECRUITER_BASE_URL")
	if len(baseUrl) == 0 {
		panic("ZIP_RECRUITER_BASE_URL env variable not set")
	}
	return &ZipClient{BaseUrl: baseUrl, ApiKey: apiKey}
}

func (z *ZipClient) Get(request ZipRequest) (*ZipResponse, error) {
	url := fmt.Sprintf(
		"%s?search=%s&location=%s&radius_miles=%d&days_ago=%d&jobs_per_page=%d&page=%d&refined_salary=%d&api_key=%s",
		z.BaseUrl,
		urlUtils.QueryEscape(request.Search),
		urlUtils.QueryEscape(request.Location),
		request.RadiusMiles,
		request.Page,
		request.JobsPerPage,
		request.DaysAgo,
		request.RefineSalary,
		z.ApiKey)
	req, errNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if errNewRequest != nil {
		log.Printf("Error: %s", errNewRequest.Error())
		return nil, errNewRequest
	}
	res, errCall := http.DefaultClient.Do(req)
	if errCall != nil {
		log.Printf("Error: %s", errCall.Error())
		return nil, errCall
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}
	if res.StatusCode != 200 {
		log.Printf("Error: %s", unauthorizedAccess)
		return nil, errors.New(unauthorizedAccess)
	}
	var zipResponse ZipResponse
	parserErr := json.Unmarshal(body, &zipResponse)
	if parserErr != nil {
		log.Printf("Error: %s", parserErr)
		return nil, errors.New(parserError)
	}
	return &zipResponse, nil
}
