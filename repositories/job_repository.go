package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"job-application/dtos"
	helpercontext "job-application/helpers/helper_context"
	"job-application/models"
)

type JobRepository interface {
	GetAllJobForApplicantRepository(context.Context, string) ([]models.Job, error)
	GetAllJobForAdminRepository(context.Context, string) ([]models.JobWithListApplicant, error)
	GetJobIdRepository(context.Context, int64) (*models.Job, error)
	PutQuotaJob(context.Context, int64, int64) error
	PostJobRepository(context.Context, dtos.RequestCreateJob) (*models.Job, error)
	PatchCloseJob(ctx context.Context, jobId int64) (*models.Job, error)
}

type JobRepositoryImplementation struct {
	db *sql.DB
}

func NewJobRepositoryImplementation(db *sql.DB) *JobRepositoryImplementation {
	return &JobRepositoryImplementation{
		db: db,
	}
}

func (jr *JobRepositoryImplementation) GetAllJobForAdminRepository(ctx context.Context, query string) ([]models.JobWithListApplicant, error) {
	searchKey := "%" + query + "%"
	sql := `
		SELECT
		*
		FROM jobs
		WHERE title ILIKE $1
		ORDER BY id ASC;
	`
	rows, err := jr.db.QueryContext(ctx, sql, searchKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	jobs := []models.JobWithListApplicant{}
	for rows.Next() {
		var job models.JobWithListApplicant
		err := rows.Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sql := `
			SELECT 
				u.id,
				u.name,
				u.email,
				u.password,
				u.role,
				u.created_at,
				u.updated_at,
				u.deleted_at
			FROM users u  
			RIGHT JOIN jobapplicants j2 ON j2.user_id = u.id
			WHERE u.deleted_at IS NULL AND j2.job_id = $1
			ORDER BY j2.job_id ASC;
		`
		rowsUser, err := jr.db.QueryContext(ctx, sql, job.ID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer rowsUser.Close()
		var users []models.User
		for rowsUser.Next() {
			var user models.User
			err := rowsUser.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.Role,
				&user.CreatedAt,
				&user.UpdatedAt,
				&user.DeleteAt,
			)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			users = append(users, user)
		}
		job.Applicants = users
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (jr *JobRepositoryImplementation) GetAllJobForApplicantRepository(ctx context.Context, query string) ([]models.Job, error) {
	serchKey := "%" + query + "%"
	sql := `
		SELECT 
		*
		FROM jobs 
		WHERE is_open = TRUE AND deleted_at IS NULL AND quota != 0 AND exp_date > CURRENT_TIMESTAMP AND title ILIKE $1
		ORDER BY id ASC;
	`
	rows, err := jr.db.QueryContext(ctx, sql, serchKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	jobs := []models.Job{}
	for rows.Next() {
		var job models.Job
		err := rows.Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (jr *JobRepositoryImplementation) GetJobIdRepository(ctx context.Context, jobId int64) (*models.Job, error) {
	sql := `
		SELECT
		*
		FROM jobs
		WHERE id = $1 AND deleted_at IS NULL AND is_open = TRUE AND quota != 0 AND exp_date > CURRENT_TIMESTAMP;
	`
	var job models.Job
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, jobId).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
	} else {
		err = jr.db.QueryRowContext(ctx, sql, jobId).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &job, nil
}

func (jr *JobRepositoryImplementation) PutQuotaJob(ctx context.Context, newQuota int64, jobId int64) error {
	sql := `
		UPDATE jobs SET 
		quota = $1,
		updated_at = NOW()
		WHERE id = $2;
	`
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		_, err := txFromCtx.ExecContext(ctx, sql, newQuota, jobId)
		return err
	}
	_, err := jr.db.ExecContext(ctx, sql, newQuota, jobId)
	return err
}

func (jr *JobRepositoryImplementation) PatchCloseJob(ctx context.Context, jobId int64) (*models.Job, error) {
	sql := `
		UPDATE jobs SET 
		is_open = FALSE,
		updated_at = NOW()
		WHERE id = $1
		RETURNING *;
	`
	var job models.Job
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, jobId).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &job, nil
	} else {
		err := jr.db.QueryRowContext(ctx, sql, jobId).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &job, nil
}

// func (jr *JobRepositoryImplementation) PatchQuotaJob(ctx context.Context, jobId int64, reqBody ) (*models.Job, error) {
// 	sql := `
// 		UPDATE jobs SET
// 		quota = $1,
// 		updated_at = NOW()
// 		WHERE id = $2
// 		RETURNING *;
// 	`
// 	var job models.Job
// 	txFromCtx := helpercontext.GetTx(ctx)
// 	if txFromCtx != nil {
// 		err := txFromCtx.QueryRowContext(ctx, sql, jobId).Scan(
// 			&job.ID,
// 			&job.Title,
// 			&job.Company,
// 			&job.IsOpen,
// 			&job.Quota,
// 			&job.ExpDate,
// 			&job.CreatedAt,
// 			&job.UpdatedAt,
// 			&job.DeleteAt,
// 		)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil, err
// 		}
// 		return &job, nil
// 	} else {
// 		err := jr.db.QueryRowContext(ctx, sql, jobId).Scan(
// 			&job.ID,
// 			&job.Title,
// 			&job.Company,
// 			&job.IsOpen,
// 			&job.Quota,
// 			&job.ExpDate,
// 			&job.CreatedAt,
// 			&job.UpdatedAt,
// 			&job.DeleteAt,
// 		)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil, err
// 		}
// 	}
// 	return &job, nil
// }

func (jr *JobRepositoryImplementation) PostJobRepository(ctx context.Context, reqBody dtos.RequestCreateJob) (*models.Job, error) {
	sql := `
		INSERT INTO jobs (title, company, is_open, quota, exp_date, created_at, updated_at) VALUES 
		($1, $2, $3, $4, $5, NOW(), NOW()) 
		RETURNING *;
	`
	var job models.Job
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err := txFromCtx.QueryRowContext(ctx, sql, reqBody.Title, reqBody.Company, reqBody.IsOpen, reqBody.Quota, reqBody.ExpDate).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := jr.db.QueryRowContext(ctx, sql, reqBody.Title, reqBody.Company, reqBody.IsOpen, reqBody.Quota, reqBody.ExpDate).Scan(
			&job.ID,
			&job.Title,
			&job.Company,
			&job.IsOpen,
			&job.Quota,
			&job.ExpDate,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return &job, nil
}
