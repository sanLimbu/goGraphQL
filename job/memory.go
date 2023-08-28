package job

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	jobs map[string][]Job
	sync.Mutex
}

// NewMemoryRepository initializes a memory with mock data
func NewMemoryRepository() *InMemoryRepository {

	jobs := make(map[string][]Job)
	jobs["1"] = []Job{
		{
			ID:         "1",
			EmployeeID: "1",
			Company:    "Google",
			Title:      "Logo",
			Start:      "2021-01-01",
			End:        "",
		},
	}
	jobs["2"] = []Job{
		{
			ID:         "2",
			EmployeeID: "2",
			Company:    "Google",
			Title:      "Janitor",
			Start:      "2021-05-03",
			End:        "",
		}, {
			ID:         "3",
			EmployeeID: "2",
			Company:    "Microsoft",
			Title:      "Janitor",
			Start:      "1980-03-04",
			End:        "2021-05-02",
		},
	}

	return &InMemoryRepository{
		jobs: jobs,
	}
}

// GetJobs returns all jobs for a certain Employee

func (im *InMemoryRepository) GetJobs(employeeID string, companyName string) ([]Job, error) {
	if jobs, ok := im.jobs[employeeID]; ok {

		filtered := make([]Job, 0)
		for _, job := range jobs {
			if (job.Company == companyName) || companyName == "" {
				filtered = append(filtered, job)
			}
		}

		return filtered, nil
	}
	return nil, errors.New("no such employee exists")
}

// GetJob will return a job based on the ID
func (im *InMemoryRepository) GetJob(employeeID, jobID string) (Job, error) {
	if jobs, ok := im.jobs[employeeID]; ok {
		for _, job := range jobs {
			if job.ID == jobID {
				return job, nil
			}
		}
		return Job{}, errors.New("no such job exists for that employee")
	}
	return Job{}, errors.New("no such employee exist")
}

// Update
func (im *InMemoryRepository) Update(j Job) (Job, error) {
	im.Lock()
	defer im.Unlock()

	if jobs, ok := im.jobs[j.EmployeeID]; ok {
		for i, job := range jobs {
			if job.ID == j.ID {
				// Replace the whole instance by index
				im.jobs[j.EmployeeID][i] = j
				return j, nil
			}
		}
	}
	return Job{}, errors.New("no such employee exists")
}
