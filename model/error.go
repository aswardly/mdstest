package model

//ValidationError represents an error resulting from failed validation
type ValidationError struct {
	ErrorField	string		//field name that failed validation
	ErrorMsg	string		//error message
}

//Error returns the error message as string (satisfies error interface)
func (ve ValidationError) Error() string {
	return ve.ErrorMsg
}

