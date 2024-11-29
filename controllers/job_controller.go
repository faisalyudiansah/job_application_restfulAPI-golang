package controllers

import (
	"io"
	"net/http"
	"strconv"

	"job-application/apperrors"
	"job-application/constants"
	"job-application/dtos"
	"job-application/helpers"
	helpercontext "job-application/helpers/helper_context"
	"job-application/services"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	JobService services.JobService
}

func NewJobController(js *services.JobServiceImplementation) *JobController {
	return &JobController{
		JobService: js,
	}
}

func (jc *JobController) GetAllJobListController(c *gin.Context) {
	query := c.Query("title")
	jobs, err := jc.JobService.GetListJobService(c, query, helpercontext.GetValueRoleFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterListJob(jobs, constants.Ok))
}

func (jc *JobController) PostApplyJobController(c *gin.Context) {
	jobId, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		c.Error(apperrors.ErrJobIdNotValid)
		return
	}
	jobs, err := jc.JobService.PostApplyJobService(c, helpercontext.GetValueRoleFromToken(c), helpercontext.GetValueUserIdFromToken(c), int64(jobId))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterApplyJob(jobs, constants.SuccessApply))
}

func (jc *JobController) PostCreateJobController(c *gin.Context) {
	var reqBody dtos.RequestCreateJob
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	jobs, err := jc.JobService.PostCreateJobService(c, reqBody, helpercontext.GetValueRoleFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterAJob(*jobs, constants.SuccessApply))
}

func (jc *JobController) PatchCloseJobController(c *gin.Context) {
	jobId, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		c.Error(apperrors.ErrJobIdNotValid)
		return
	}
	jobs, err := jc.JobService.PatchJobCloseService(c, int64(jobId), helpercontext.GetValueRoleFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterAJob(*jobs, constants.SuccessCloseJob))
}
