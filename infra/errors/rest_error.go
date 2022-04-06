package errors

import (
	"core/infra/logger"
	"errors"
	"net/http"
)

var (
	ErrDuplicateEntry            uint16 = 1062
	ErrInvalidEmail                     = NewError("invalid email")
	ErrInvalidPhone                     = NewError("invalid phone no")
	ErrInvalidPassword                  = NewError("invalid password")
	ErrUserRolePermissions              = NewError("failed to fetch role permissions")
	ErrCreateJwt                        = NewError("failed to create JWT token")
	ErrAccessTokenSign                  = NewError("failed to sign access_token")
	ErrRefreshTokenSign                 = NewError("failed to sign refresh_token")
	ErrStoreTokenUuid                   = NewError("failed to store token uuid")
	ErrUpdateLastLogin                  = NewError("failed to update last login")
	ErrNoContextUser                    = NewError("failed to get user from context")
	ErrInvalidRefreshToken              = NewError("invalid refresh_token")
	ErrInvalidAccessToken               = NewError("invalid access_token")
	ErrInvalidPasswordResetToken        = NewError("invalid reset_token")
	ErrInvalidConfirmationlToken        = NewError("invalid confirmationl_token")
	ErrInvalidRefreshUuid               = NewError("invalid refresh_uuid")
	ErrInvalidAccessUuid                = NewError("invalid refresh_uuid")
	ErrInvalidJwtSigningMethod          = NewError("invalid signing method while parsing jwt")
	ErrParseJwt                         = NewError("failed to parse JWT token")
	ErrDeleteOldTokenUuid               = NewError("failed to delete old token uuids")
	ErrSendingEmail                     = NewError("failed to send email")
	ErrNotAdmin                         = NewError("not admin")
	ErrNotSuperAdmin                    = NewError("not super admin")
	ErrEmptyRedisKeyValue               = NewError("empty redis key or value")
	ErrInvalidJsonBody                  = NewError("invalid json body")
	ErrPhoneOrEmailExists               = "email or phone number already exists"
	ErrProductNameExists                = "product name already exists"
	ErrUserNameNotUnique                = "username already exists!, try another"
	ErrCompanyNameNotUnique             = "company name already exists!, try another"
	ErrPhoneNoIsUnique                  = "phone no already exists!, try another"
	ErrEmailIsUnique                    = "email already exists!, try another"
	ErrPhoneNameNotUnique               = "phone number already exists!, try another"
	ErrSomethingWentWrong               = "something went wrong"
	ErrRecordNotFound                   = "record not found"
	ErrRecordNotvalid                   = "invalid parameters, check name, email or username "
	ErrCreateUserNotvalid               = "invalid parameters, check display name, email or password"
	ErrCheckParamBodyHeader             = "check header, params, body"
	ErrNoLoggedInUser                   = "failed to get logged in user"
)

const (
	ErrSingleFileValidation   = "Please upload one image"
	ErrMultipleFileValidation = "No images received"
)

type RestErr struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Status  int    `json:"status"`
}

func (err *RestErr) Error() string {
	return err.Message
}

func NewError(msg string) error {
	return errors.New(msg)
}

func As(err error, target interface{}) bool {
	return errors.As(err, &target)
}

func NewInternalServerError(err error) *RestErr {
	restErr := &RestErr{
		Message: ErrSomethingWentWrong,
		Status:  http.StatusInternalServerError,
	}

	if err != nil {
		logger.ErrorAsJson(err.Error(), err)
		restErr.Detail = err.Error()
	} else {
		logger.ErrorAsJson(restErr.Error(), restErr)
	}

	return restErr
}

func NewBadRequestError(message string, err error) *RestErr {
	restErr := &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
	}

	if err != nil {
		logger.ErrorAsJson(err.Error(), err)
		restErr.Detail = err.Error()
	} else {
		logger.ErrorAsJson(restErr.Error(), restErr)
	}

	return restErr
}

func NewNotFoundError(message string, err error) *RestErr {
	restErr := &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
	}

	if err != nil {
		logger.ErrorAsJson(err.Error(), err)
		restErr.Detail = err.Error()
	} else {
		logger.ErrorAsJson(restErr.Error(), restErr)
	}

	return restErr
}

func NewAlreadyExistError(message string, err error) *RestErr {
	restErr := &RestErr{
		Message: message,
		Status:  http.StatusConflict,
	}

	if err != nil {
		logger.ErrorAsJson(err.Error(), err)
		restErr.Detail = err.Error()
	} else {
		logger.ErrorAsJson(restErr.Error(), restErr)
	}

	return restErr
}

func NewUnauthorizedError(message string, err error) *RestErr {
	restErr := &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
	}

	if err != nil {
		logger.ErrorAsJson(err.Error(), err)
		restErr.Detail = err.Error()
	} else {
		logger.ErrorAsJson(restErr.Error(), restErr)
	}

	return restErr
}
