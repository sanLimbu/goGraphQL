package gopher

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sanLimbu/gopheragency/job"
)

type Resolver interface {

	// ResolveGophers should return a list of all gophers in the repository
	ResolveGophers(p graphql.ResolveParams) (interface{}, error)
	//ResolveGopher is used to respond to single queries for gophers
	ResolveGopher(p graphql.ResolveParams) (interface{}, error)
	//ResolveJobs is used to find jobs
	ResolveJobs(p graphql.ResolveParams) (interface{}, error)
}

// GopherService is the service that holds all repositories

type GopherService struct {
	gophers Repository
	jobs    job.Repository
}

// NewService is a factory that creates a new GopherService
func NewService(repo Repository, jobRepo job.Repository) GopherService {
	return GopherService{
		gophers: repo,
		jobs:    jobRepo,
	}
}

// ResolveGophers will be used to retrieve all available Gophers

func (gs GopherService) ResolveGophers(p graphql.ResolveParams) (interface{}, error) {
	//fetch gophers from the repository
	gophers, err := gs.gophers.GetGophers()
	if err != nil {
		return nil, err
	}
	return gophers, nil
}

// ResolveJobs is used to find all jobs related to a gopher
func (gs *GopherService) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {
	//fetch source value
	g, ok := p.Source.(Gopher)
	if !ok {
		return nil, errors.New("source was not gopher")
	}

	//We extract the argument company
	company := ""
	if value, ok := p.Args["company"]; ok {
		company, ok = value.(string)
		if !ok {
			return nil, errors.New("id has to be string")
		}
	}

	//Find the jobs based on the gophers ID
	jobs, err := gs.jobs.GetJobs(g.ID, company)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// grab a string argument
func grabStringArgument(k string, args map[string]interface{}, required bool) (string, error) {
	//first check the presence of arg
	if value, ok := args[k]; ok {
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}

func (gs *GopherService) MutateJobs(p graphql.ResolveParams) (interface{}, error) {
	employee, err := grabStringArgument("employeeid", p.Args, true)
	if err != nil {
		return nil, err
	}
	jobid, err := grabStringArgument("jobid", p.Args, true)
	if err != nil {
		return nil, err
	}
	start, err := grabStringArgument("start", p.Args, false)
	if err != nil {
		return nil, err
	}
	end, err := grabStringArgument("end", p.Args, false)
	if err != nil {
		return nil, err
	}

	// Get the job
	job, err := gs.jobs.GetJob(employee, jobid)
	if err != nil {
		return nil, err
	}

	// Modify start and end date if they are set
	if start != "" {
		job.Start = start
	}

	if end != "" {
		job.End = end
	}
	return gs.jobs.Update(job)
}
