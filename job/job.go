package job

// Repository is used to specify whats needed to fulfill the job storage requirements

type Repository interface {
	//GetJobs will search for all the jobs related to empployeeID
	GetJobs(empployeeID string, company string) ([]Job, error)

	GetJob(employeeID, jobid string) (Job, error)

	Update(Job) (Job, error)
}

type Job struct {
	ID string `json:"id"`
	// EmployeeID is the employee related to the job
	EmployeeID string `json:"employeeID"`
	Company    string `json:"company"`
	Title      string `json:"title"`
	// Start is when the job started
	Start string `json:"start"`
	// End is when the employment ended
	End string `json:"end"`
}
