package exception

const CodeRetSuccess = 0

const (
	CodeInternalError = iota + 1
	CodeQueryFailed
	CodeUnableConnect
	CodeForbidden
	CodeUnauthorized
	CodeNoPermission
	CodeAccountDisabled
)

const (
	CodeInvalidParams = iota + 101
	CodeConvertFailed
	CodeDataNotExist
	CodeDataAlreadyExist
	CodeDataCantSet
	CodeOperateTooFast
	CodeCallApiFailed
	CodeUploadFailed
	CodeUploadFileTypeError
	CodeOtpVerifyFailed
	CodeOtpSendFailed
)

const (
	CodeMenuParentNotExist = iota + 1001
	CodeMenuPermsInUse
	CodeUserPasswordErr
)

var Desces = map[int]string{
	CodeRetSuccess:      "success",
	CodeInternalError:   "server internal error",
	CodeQueryFailed:     "data query failed",
	CodeUnableConnect:   "unable to connect to server",
	CodeForbidden:       "access denied",
	CodeUnauthorized:    "unauthorized",
	CodeNoPermission:    "no permission",
	CodeAccountDisabled: "account is disabled",

	CodeOtpVerifyFailed: "verify otp failed",
	CodeOtpSendFailed:   "send otp failed",

	CodeInvalidParams:       "invalid parameter",
	CodeConvertFailed:       "convert data failed",
	CodeDataNotExist:        "data not exist",
	CodeDataAlreadyExist:    "data already exist",
	CodeDataCantSet:         "set data failed",
	CodeOperateTooFast:      "operate too fast, please try again later",
	CodeCallApiFailed:       "call API interface failed",
	CodeUploadFailed:        "upload file failed",
	CodeUploadFileTypeError: "upload file type error",

	CodeMenuParentNotExist: "parent menu node not exist",
	CodeMenuPermsInUse:     "menu perms already in use",
	CodeUserPasswordErr:    "name or password error",
}
