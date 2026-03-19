package idgen

import "github.com/segmentio/ksuid"


// NewUserID generates a new KSUID for users.user_id
func NewUserID() string {
	return ksuid.New().String()
}

// NewGroupID generates a new KSUID for groups.group_id
func NewGroupID() string {
	return ksuid.New().String()
}

// NewEndpointID generates a new KSUID for endpoint.endpoint_id
func NewEndpointID() string {
	return ksuid.New().String()
}

// NewGroupEndpointID generates a new KSUID for group_endpoint.id
func NewGroupEndpointID() string {
	return ksuid.New().String()
}