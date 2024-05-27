package http

const RequestIdContextKey = "RequestID"
const UserIdContextKey = "UserID"
const SellerUserIdContextKey = "SellerUserID"
const RequestIdHeaderKey = "X-Request-ID"

func GetStandartSuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
