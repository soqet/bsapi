package bsapi

type ClientError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (ce ClientError) Error() string {
	return ce.Reason
}
