package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "job-application/helpers/helper_context"
	"job-application/models"
)

type JobApplicantRepository interface {
	IsUserAlreadyApplyToThatJobRepository(context.Context, int64, int64) (*models.JobApplicant, error)
	PostJobApplicantRepository(context.Context, int64, int64) (*models.JobApplicant, error)
}

type JobApplicantRepositoryImplementation struct {
	db *sql.DB
}

func NewJobApplicantRepositoryImplementation(db *sql.DB) *JobApplicantRepositoryImplementation {
	return &JobApplicantRepositoryImplementation{
		db: db,
	}
}

func (jr *JobApplicantRepositoryImplementation) IsUserAlreadyApplyToThatJobRepository(ctx context.Context, jobId int64, userId int64) (*models.JobApplicant, error) {
	sql := `
		SELECT
		*
		FROM JobApplicants
		WHERE job_id = $1 AND user_id = $2 AND deleted_at IS NULL;
	`
	var job models.JobApplicant
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, jobId, userId).Scan(
			&job.ID,
			&job.UserId,
			&job.JobId,
			&job.Status,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
	} else {
		jr.db.QueryRowContext(ctx, sql, jobId, userId).Scan(
			&job.ID,
			&job.UserId,
			&job.JobId,
			&job.Status,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
	}
	return &job, nil
}

func (jr *JobApplicantRepositoryImplementation) PostJobApplicantRepository(ctx context.Context, jobId int64, userId int64) (*models.JobApplicant, error) {
	sql := `
		INSERT INTO JobApplicants (user_id, job_id, status, created_at, updated_at)
		VALUES
		($1, $2, 'applied', NOW(), NOW())
		RETURNING *;
	`
	var job models.JobApplicant
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId, jobId).Scan(
			&job.ID,
			&job.UserId,
			&job.JobId,
			&job.Status,
			&job.CreatedAt,
			&job.UpdatedAt,
			&job.DeleteAt,
		)
	} else {
		err = jr.db.QueryRowContext(ctx, sql, jobId, userId).Scan(
			&job.ID,
			&job.UserId,
			&job.JobId,
			&job.Status,
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
