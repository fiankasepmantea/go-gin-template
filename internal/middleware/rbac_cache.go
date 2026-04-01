package middleware

var RBACCache = make(map[string]bool)

func BuildKey(userID string, endpointID string) string {
	return userID + ":" + endpointID
}