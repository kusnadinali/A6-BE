package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type guidelinesActive struct {
	GuidelinesName string
	GuidelinesDesc string
	GuidelinesType string
	GuidelinesLink string
}
