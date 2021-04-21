package response

type Error struct {
	HttpStatusCode int                   `json:"httpStatusCode"`
	ErrorCode      string                `json:"errorCode,omitempty"`
	Message        string                `json:"message"`
	Info           string                `json:"info,omitempty"`
	AdditionalInfo []AdditionalInfoError `json:"additionalInfo,omitempty"`
}
