package apierrors

var Auth = map[string]ApiErrorCodeStatus{
	"email_already_taken": {
		Code:   "auth-0001",
		Status: "email_already_taken",
	},
	"username_already_taken": {
		Code:   "auth-0002",
		Status: "username_already_taken",
	},
}
