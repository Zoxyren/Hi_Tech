package errorHandling

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
func NewError(status int, message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: message,
	}
}

var (
	ErrItemWithIDNotFound   = NewError(http.StatusNotFound, "Item with given ID not found")
	ErrItemsNotFound        = NewError(http.StatusNotFound, "Items not found")
	ErrItemNotFound         = NewError(http.StatusNotFound, "item not found")
	ErrInvalidCredentials   = NewError(http.StatusUnauthorized, "invalid credentials")
	ErrInternalServer       = NewError(http.StatusInternalServerError, "internal server error")
	ErrProductAlreadyExists = NewError(http.StatusConflict, "product already exists")
	ErrInsertingStatement   = NewError(http.StatusInternalServerError, "error inserting statement")
	ErrorAddingProduct      = NewError(http.StatusInternalServerError, "error adding product")
)
