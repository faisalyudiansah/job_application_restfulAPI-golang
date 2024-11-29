package constants

const (
	ISE                = "internal server error"
	InvalidAccessToken = "invalid access token"
	Unauthorization    = "unauthorization"
	UrlNotFound        = "url not found"
	RequestBodyInvalid = "request body invalid or missing"
)

const (
	UserInvalidEmailPassword = "invalid email / password"
	UserEmailAlreadyExists   = "email already exists"
	UserFailedRegister       = "there was an error in the register process, try again"
)

const (
	JobIdNotExists            = "job id is not exists"
	FailedApplyJob            = "there was an error when apply a job process, try again"
	FailedCreateJob           = "there was an error when create a job process, try again"
	JobIdNotValid             = "job id is not valid"
	UserAlreadyApplyToThisJob = "user already apply to this job"
	QuotaInvalid              = "quota is invalid"
	FailedPatchIsOpen         = "something wrong when change status job, try again"
)
