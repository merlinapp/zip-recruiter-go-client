package jobs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	urlUtils "net/url"
	"os"
	"strconv"
)

const (
	parserError        = "could not parse error"
	unauthorizedAccess = "unauthorized access"
	urlParseError      = "could not parse base url"
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
	u, err := z.buildQueryParams(request)
	if err != nil {
		return nil, err
	}
	req, errNewRequest := http.NewRequest(http.MethodGet, u.String(), nil)
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

func (z *ZipClient) buildQueryParams(request ZipRequest) (*urlUtils.URL, error) {
	u, err := urlUtils.Parse(z.BaseUrl)
	if err != nil {
		log.Printf("Error: %s", urlParseError)
		return nil, errors.New(urlParseError)
	}
	queryParams := u.Query()
	queryParams.Set("search", request.Search)
	queryParams.Set("location", request.Location)
	queryParams.Set("radius_miles", strconv.Itoa(int(request.RadiusMiles)))
	queryParams.Set("days_ago", strconv.Itoa(int(request.DaysAgo)))
	queryParams.Set("jobs_per_page", strconv.Itoa(int(request.JobsPerPage)))
	queryParams.Set("page", strconv.Itoa(int(request.Page)))
	queryParams.Set("refine_by_salary", strconv.Itoa(int(request.RefineSalary)))
	queryParams.Set("api_key", z.ApiKey)
	u.RawQuery = queryParams.Encode()
	return u, err
}
