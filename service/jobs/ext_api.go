package jobs

import "context"

// GetRun retrieves a run based on the provided request.
// It handles pagination if the run contains multiple iterations or tasks.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	run, err := a.jobsImpl.GetRun(ctx, request)
	if err != nil {
		return nil, err
	}

	// When querying a Job run, a page token is returned when there are more than 100 tasks. No iterations are defined for a Job run. Therefore, the next page in the response only includes the next page of tasks.
	// When querying a ForEach task run, a page token is returned when there are more than 100 iterations. Only a single task is returned, corresponding to the ForEach task itself. Therefore, the client only reads the iterations from the next page and not the tasks.
	isPaginatingIterations := len(run.Iterations) > 0

	// runs/get response includes next_page_token as long as there are more pages to fetch.
	for run.NextPageToken != "" {
		request.PageToken = run.NextPageToken
		nextRun, err := a.jobsImpl.GetRun(ctx, request)
		if err != nil {
			return nil, err
		}

		if isPaginatingIterations {
			run.Iterations = append(run.Iterations, nextRun.Iterations...)
		} else {
			run.Tasks = append(run.Tasks, nextRun.Tasks...)
		}
		// Each new page of runs/get response includes the next page of the job_clusters, job_parameters, and repair history.
		run.JobClusters = append(run.JobClusters, nextRun.JobClusters...)
		run.JobParameters = append(run.JobParameters, nextRun.JobParameters...)
		run.RepairHistory = append(run.RepairHistory, nextRun.RepairHistory...)
		run.NextPageToken = nextRun.NextPageToken
	}

	return run, nil
}

// Get retrieves a job based on the provided request.
// It handles pagination if the job contains multiple tasks, job_clusters, job_parameters or environments.
func (a *JobsAPI) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	job, err := a.jobsImpl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	// jobs/get response includes next_page_token as long as there are more pages to fetch.
	for job.NextPageToken != "" {
		request.PageToken = job.NextPageToken
		nextJob, err := a.jobsImpl.Get(ctx, request)
		if err != nil {
			return nil, err
		}

		// Each new page of jobs/get response includes the next page of the tasks, job_clusters, job_parameters, and environments.
		job.Settings.Tasks = append(job.Settings.Tasks, nextJob.Settings.Tasks...)
		job.Settings.JobClusters = append(job.Settings.JobClusters, nextJob.Settings.JobClusters...)
		job.Settings.Parameters = append(job.Settings.Parameters, nextJob.Settings.Parameters...)
		job.Settings.Environments = append(job.Settings.Environments, nextJob.Settings.Environments...)
		job.NextPageToken = nextJob.NextPageToken
	}

	return job, nil
}
