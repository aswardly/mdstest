package response

const ResponseCodeSuccess = "OK"
const ResponseCodeError = "ERR"

//GenericResponse represents a generic response to return to the echo client
type GenericResponse struct {
	ResponseCode 	string		`json:"response_code"`
	ResponseMessage string		`json:"response_message"`
	Data 			interface{}	`json:"data"`
}