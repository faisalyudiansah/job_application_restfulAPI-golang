package apperrors

import (
	"errors"

	"job-application/constants"
)

var (
	ErrISE                = errors.New(constants.ISE)
	ErrInvalidAccessToken = errors.New(constants.InvalidAccessToken)
	ErrUnauthorization    = errors.New(constants.Unauthorization)
	ErrUrlNotFound        = errors.New(constants.UrlNotFound)
	ErrRequestBodyInvalid = errors.New(constants.RequestBodyInvalid)
)

var (
	ErrUserEmailAlreadyExists   = errors.New(constants.UserEmailAlreadyExists)
	ErrUserFailedRegister       = errors.New(constants.UserFailedRegister)
	ErrUserInvalidEmailPassword = errors.New(constants.UserInvalidEmailPassword)
)

var (
	ErrJobIdNotExists            = errors.New(constants.JobIdNotExists)
	ErrFailedPatchIsOpen         = errors.New(constants.FailedPatchIsOpen)
	ErrFailedCreateJob           = errors.New(constants.FailedCreateJob)
	ErrFailedApplyJob            = errors.New(constants.FailedApplyJob)
	ErrUserAlreadyApplyToThisJob = errors.New(constants.UserAlreadyApplyToThisJob)
	ErrJobIdNotValid             = errors.New(constants.JobIdNotValid)
	ErrQuotaInvalid              = errors.New(constants.QuotaInvalid)
)
