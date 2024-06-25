package excep

const (
	EXCEP_UNAUTHORIZED          string = "UNAUTHORIZED"
	EXCEP_FORBIDDEN             string = "FORBIDDEN"
	EXCEP_INVALID_PAYLOAD       string = "INVALID_PAYLOAD"
	EXCEP_INTERNAL_SERVER_ERROR string = "INTERNAL_SERVER_ERROR"

	// CUSTOM EXCEPTIONS
	EXCEP_USER_NOT_FOUND string = "USER_NOT_FOUND"
	EXCEP_EMAIL_EXISTS   string = "EMAIL_EXISTS"
)

type DomainExcep struct {
	Type    string `json:"type"`
	Details string `json:"details"`
}

func (d DomainExcep) Error() string {
	return d.Details
}

func NewUnauthorized(details string) DomainExcep {
	if details == "" {
		details = "Unauthorized"
	}

	return DomainExcep{
		Type:    EXCEP_UNAUTHORIZED,
		Details: details,
	}
}

func NewForbidden(details string) DomainExcep {
	if details == "" {
		details = "Forbidden"
	}

	return DomainExcep{
		Type:    EXCEP_FORBIDDEN,
		Details: details,
	}
}

func NewInvalidPayload(details string) DomainExcep {
	if details == "" {
		details = "Invalid payload"
	}

	return DomainExcep{
		Type:    EXCEP_INVALID_PAYLOAD,
		Details: details,
	}
}

func NewInternalServerError(details string) DomainExcep {
	if details == "" {
		details = "Internal server error"
	}

	return DomainExcep{
		Type:    EXCEP_INTERNAL_SERVER_ERROR,
		Details: details,
	}
}
