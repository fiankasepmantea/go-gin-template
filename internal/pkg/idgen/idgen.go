package idgen

import "github.com/segmentio/ksuid"

func newKSUID() string {
	return ksuid.New().String()
}

func NewID() string          { return newKSUID() }
func NewUserID() string          { return newKSUID() }
func NewUserTokenID() string     { return newKSUID() }
func NewGroupID() string         { return newKSUID() }
func NewEndpointID() string      { return newKSUID() }
func NewGroupEndpointID() string { return newKSUID() }
func NewRefreshToken() string    { return newKSUID() }