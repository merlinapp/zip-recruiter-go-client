package jobs

type ZipRequest struct {
	Search       string
	Location     string
	RadiusMiles  string
	Page         string
	JobsPerPage  string
	DaysAgo      string
	RefineSalary string
}

type HiringCompanyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type JobResponse struct {
	Source             string                `json:"source"`
	ID                 string                `json:"id"`
	Name               string                `json:"name"`
	Snippet            string                `json:"snippet"`
	Category           string                `json:"category"`
	PostedTime         string                `json:"posted_time"`
	PostedTimeFriendly string                `json:"posted_time_friendly"`
	Url                string                `json:"url"`
	Location           string                `json:"location"`
	City               string                `json:"city"`
	State              string                `json:"state"`
	Country            string                `json:"country"`
	SalarySource       string                `json:"salary_source"`
	SalaryInterval     string                `json:"salary_interval"`
	SalaryMax          float64               `json:"salary_max"`
	SalaryMaxAnnual    float64               `json:"salary_max_annual"`
	SalaryMin          float64               `json:"salary_min"`
	SalaryMinAnnual    float64               `json:"salary_min_annual"`
	IndustryName       string                `json:"industry_name"`
	HiringCompany      HiringCompanyResponse `json:"hiring_company"`
}

type ZipResponse struct {
	Succeed          bool          `json:"success"`
	TotalJobs        int64         `json:"total_jobs"`
	NumPaginableJobs int64         `json:"num_paginable_jobs"`
	Jobs             []JobResponse `json:"jobs"`
}
