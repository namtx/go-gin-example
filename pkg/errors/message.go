package errors

var Messages = map[int]string{
	Success:                 "ok",
	Error:                   "fail",
	InvalidParams:           "invalid params",
	ErrorAuthCheckTokenFail: "invalid token",
	ErrorIssueAuthTokenFail: "failed to issue new token",
}

func GetMessage(errCode int) string {
	msg, ok := Messages[errCode]
	if ok {
		return msg
	}

	return Messages[Error]
}
