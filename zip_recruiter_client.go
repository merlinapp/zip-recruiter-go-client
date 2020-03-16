package zip_recruiter_go_client

type Client interface {
	Get(apiKey string,
		search string,
		location string,
		radiusMiles string,
		page string,
		jobsPerPage string,
		daysAgo string,
		refineSalary string,
	) ([]byte, int, error)
}
