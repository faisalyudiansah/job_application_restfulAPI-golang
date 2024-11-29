package helpers

import (
	"job-application/dtos"
	"job-application/models"
)

func ModelToDtoResponseRegister(dataUser models.UserAndProfile) *dtos.ResponseRegisterUserAndUserProfile {
	return &dtos.ResponseRegisterUserAndUserProfile{
		ID:        dataUser.User.ID,
		Name:      dataUser.User.Name,
		Email:     dataUser.User.Email,
		Role:      dataUser.User.Role,
		CreatedAt: dataUser.User.CreatedAt,
		UpdatedAt: dataUser.User.UpdatedAt,
		DeleteAt:  dataUser.User.DeleteAt,
		Detail:    dtos.ResponseUserProfile(dataUser.UserProfile),
	}
}

func AccessTokenToDtoResponseUserAccessToken(ac string) *dtos.ResponseAccessToken {
	return &dtos.ResponseAccessToken{
		AccessToken: ac,
	}
}

func JobsToDtoResponseListJobApplicant(jobs []models.Job) []dtos.ResponseJobApplicant {
	var dtoJobs []dtos.ResponseJobApplicant
	for _, job := range jobs {
		dtoJobs = append(dtoJobs, ModelToDtoResponseJobApplicant(job))
	}
	return dtoJobs
}

func JobsToDtoResponseListJobAdmin(jobs []models.JobWithListApplicant) []dtos.ResponseJobApplicant {
	var dtoJobs []dtos.ResponseJobApplicant
	for _, job := range jobs {
		dtoJobs = append(dtoJobs, ModelToDtoResponseJobApplicantWithListApplicant(job))
	}
	return dtoJobs
}

func ModelToDtoResponseJobApplicant(dataJob models.Job) dtos.ResponseJobApplicant {
	return dtos.ResponseJobApplicant{
		ID:        dataJob.ID,
		Title:     dataJob.Title,
		Company:   dataJob.Company,
		IsOpen:    dataJob.IsOpen,
		Quota:     dataJob.Quota,
		ExpDate:   dataJob.ExpDate,
		CreatedAt: dataJob.CreatedAt,
		UpdatedAt: dataJob.UpdatedAt,
		DeleteAt:  dataJob.DeleteAt,
	}
}

func ModelToDtoResponseCreateAJob(dataJob *models.Job) *dtos.ResponseJob {
	return &dtos.ResponseJob{
		ID:        dataJob.ID,
		Title:     dataJob.Title,
		Company:   dataJob.Company,
		IsOpen:    dataJob.IsOpen,
		Quota:     dataJob.Quota,
		ExpDate:   dataJob.ExpDate,
		CreatedAt: dataJob.CreatedAt,
		UpdatedAt: dataJob.UpdatedAt,
		DeleteAt:  dataJob.DeleteAt,
	}
}

func ModelToDtoResponseJobApplicantWithListApplicant(dataJob models.JobWithListApplicant) dtos.ResponseJobApplicant {
	return dtos.ResponseJobApplicant{
		ID:        dataJob.ID,
		Title:     dataJob.Title,
		Company:   dataJob.Company,
		IsOpen:    dataJob.IsOpen,
		Quota:     dataJob.Quota,
		ExpDate:   dataJob.ExpDate,
		CreatedAt: dataJob.CreatedAt,
		UpdatedAt: dataJob.UpdatedAt,
		DeleteAt:  dataJob.DeleteAt,
		Applicant: &dataJob.Applicants,
	}
}

func ModelToDtoResponseApply(data models.JobApplicantWithDataUser) *dtos.ResponseApply {
	return &dtos.ResponseApply{
		ID:          data.ID,
		UserId:      data.UserId,
		JobId:       data.JobId,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeleteAt:    data.DeleteAt,
		UserProfile: dtos.ResponseUserProfile(data.UserProfile),
	}
}
