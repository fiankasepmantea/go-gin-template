package groupendpoint

import "github.com/fiankasepman/go-gin-template/internal/pkg/idgen"

type GroupEndpoint struct {
	ID         string `gorm:"column:id;primaryKey"`
	GroupID    string `gorm:"column:group_id"`
	EndpointID string `gorm:"column:endpoint_id"`
}

func NewGroupEndpointID() string {
	return idgen.NewGroupEndpointID()
}

func (GroupEndpoint) TableName() string {
	return "group_endpoint"
}