package services

import (
	"context"

	"job-application/apperrors"
	"job-application/dtos"
	"job-application/helpers"
	"job-application/models"
	"job-application/repositories"
)

type JobService interface {
	GetListJobService(context.Context, string, string) ([]dtos.ResponseJobApplicant, error)
	PostApplyJobService(context.Context, string, int64, int64) (*dtos.ResponseApply, error)
	PostCreateJobService(context.Context, dtos.RequestCreateJob, string) (*dtos.ResponseJob, error)
	PatchJobCloseService(context.Context, int64, string) (*dtos.ResponseJob, error)
}

type JobServiceImplementation struct {
	UserRepository         repositories.UserRepository
	JobRepository          repositories.JobRepository
	JobApplicantRepository repositories.JobApplicantRepository
	TransactionsRepository repositories.TransactionRepository
}

func NewJobServiceImplementation(us repositories.UserRepository, js repositories.JobRepository, ja repositories.JobApplicantRepository, tx repositories.TransactionRepository) *JobServiceImplementation {
	return &JobServiceImplementation{
		UserRepository:         us,
		JobRepository:          js,
		JobApplicantRepository: ja,
		TransactionsRepository: tx,
	}
}

func (js *JobServiceImplementation) GetListJobService(ctx context.Context, query string, role string) ([]dtos.ResponseJobApplicant, error) {
	if role == "" {
		return nil, apperrors.ErrUnauthorization
	}
	if role == "applicant" {
		jobs, err := js.JobRepository.GetAllJobForApplicantRepository(ctx, query)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return helpers.JobsToDtoResponseListJobApplicant(jobs), nil
	}
	jobs, err := js.JobRepository.GetAllJobForAdminRepository(ctx, query)
	if err != nil {
		return nil, apperrors.ErrISE
	}
	return helpers.JobsToDtoResponseListJobAdmin(jobs), nil
}

func (js *JobServiceImplementation) PostApplyJobService(ctx context.Context, role string, userId int64, jobId int64) (*dtos.ResponseApply, error) {
	if role == "admin" {
		return nil, apperrors.ErrUnauthorization
	}
	result, err := js.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		findUser, err := js.UserRepository.GetUserById(cForTx, userId)
		if err != nil {
			return nil, apperrors.ErrInvalidAccessToken
		}
		findUserProfile, err := js.UserRepository.GetUserProfileById(cForTx, userId)
		if err != nil {
			return nil, apperrors.ErrInvalidAccessToken
		}
		getJob, err := js.JobRepository.GetJobIdRepository(cForTx, jobId)
		if err != nil || getJob.ID == 0 {
			return nil, apperrors.ErrJobIdNotExists
		}
		checkUserApply, err := js.JobApplicantRepository.IsUserAlreadyApplyToThatJobRepository(cForTx, getJob.ID, findUser.ID)
		if err != nil || checkUserApply.ID != 0 {
			return nil, apperrors.ErrUserAlreadyApplyToThisJob
		}
		result, err := js.JobApplicantRepository.PostJobApplicantRepository(cForTx, getJob.ID, findUser.ID)
		if err != nil {
			return nil, apperrors.ErrFailedApplyJob
		}
		err = js.JobRepository.PutQuotaJob(cForTx, int64(getJob.Quota-1), getJob.ID)
		if err != nil {
			return nil, apperrors.ErrFailedApplyJob
		}
		return helpers.FormatterToJobApplicantWithDataUser(result, findUserProfile), nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*dtos.ResponseApply), nil
}

func (js *JobServiceImplementation) PostCreateJobService(ctx context.Context, reqBody dtos.RequestCreateJob, role string) (*dtos.ResponseJob, error) {
	if role == "applicant" {
		return nil, apperrors.ErrUnauthorization
	}
	result, err := js.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		if reqBody.Quota == 0 {
			return nil, apperrors.ErrQuotaInvalid
		}
		result, err := js.JobRepository.PostJobRepository(cForTx, reqBody)
		if err != nil {
			return nil, apperrors.ErrFailedCreateJob
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	res := helpers.ModelToDtoResponseCreateAJob(result.(*models.Job))
	return res, nil
}

func (js *JobServiceImplementation) PatchJobCloseService(ctx context.Context, jobId int64, role string) (*dtos.ResponseJob, error) {
	if role == "applicant" {
		return nil, apperrors.ErrUnauthorization
	}
	result, err := js.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		getJob, err := js.JobRepository.GetJobIdRepository(cForTx, jobId)
		if err != nil || getJob.ID == 0 {
			return nil, apperrors.ErrJobIdNotExists
		}
		res, err := js.JobRepository.PatchCloseJob(cForTx, getJob.ID)
		if err != nil {
			return nil, apperrors.ErrJobIdNotExists
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ModelToDtoResponseCreateAJob(result.(*models.Job)), nil
}

func (js *JobServiceImplementation) PatchJobQuotaService(ctx context.Context, jobId int64, role string) (*dtos.ResponseJob, error) {
	if role == "applicant" {
		return nil, apperrors.ErrUnauthorization
	}
	result, err := js.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		getJob, err := js.JobRepository.GetJobIdRepository(cForTx, jobId)
		if err != nil || getJob.ID == 0 {
			return nil, apperrors.ErrJobIdNotExists
		}
		res, err := js.JobRepository.PatchCloseJob(cForTx, getJob.ID)
		if err != nil {
			return nil, apperrors.ErrJobIdNotExists
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ModelToDtoResponseCreateAJob(result.(*models.Job)), nil
}
