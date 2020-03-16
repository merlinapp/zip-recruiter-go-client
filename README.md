# zip-recruiter-go-client
This project is used for use as a library the zip recruiter api


## Installation

Set the environment variable `ZIP_RECRUITER_KEY`

Using go modules, import

```go
import (
    "github.com/merlinapp/zip-recruiter-go-client/jobs"
)
```
so in your go modules file `go.mod` you should hve the following entry:

```
require (
    github.com/merlinapp/zip-recruiter-go-client v0.0.6
)
```
## Usage

```go
    client := jobs.NewZipClient()
    searchFilers := jobs.SearchFilters{
        Search: "cashier",    
        Location: "Manhattan, NY",    
        RadiusMiles: 3,    
        Page: 1,    
        JobsPerPage: 100,    
        DaysAgo: 2,    
        RefineSalary: 35,    
        }
    response, err := client.Get(searchFilers)
    if respose.succeed {
        fmt.printf("number of jos found: %d",response.TotalJobs)
        for i, job in range(response.Jobs){
            fmt.print(job.Source)
            fmt.print(job.ID)
            fmt.print(job.Name)
            fmt.print(job.Snippet)
            fmt.print(job.Category)
            fmt.print(job.PostedTime)
            fmt.print(job.PostedTimeFriendly)
            fmt.print(job.Url)
            fmt.print(job.Location)
            fmt.print(job.City)
            fmt.print(job.State)
            fmt.print(job.Country)
            fmt.print(job.SalarySource)
            fmt.print(job.SalaryInterval)
            fmt.print(job.SalaryMax)
            fmt.print(job.SalaryMaxAnnual)
            fmt.print(job.SalaryMin)
            fmt.print(job.SalaryMinAnnual)
            fmt.print(job.IndustryName)
            fmt.print(job.HiringCompany.ID)
            fmt.print(job.HiringCompany.Name)
            fmt.print(job.HiringCompany.Url)
            fmt.print(job.HiringCompany.Description)
        }       
    }    
```
