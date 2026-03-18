package cli


func modelTemplate(name string) string {

	structName := toPascalCase(name)

	return `package ` + name + `

import "github.com/fiankasepman/go-gin-template/internal/base"

type ` + structName + ` struct {
	base.BaseModel
	Name string
}
`
}

func repoTemplate(name string) string {

	structName := toPascalCase(name)

	return `package ` + name + `

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[` + structName + `]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[` + structName + `]{DB: db},
	}
}
`
}

func serviceTemplate(name string) string {

	return `package ` + name + `

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
`
}

func handlerTemplate(name string) string {

	structName := toPascalCase(name)

	return `package ` + name + `

import (
	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/base"
)

type Handler struct {
	base.BaseHandler
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *gin.Context) {

	var data []` + structName + `

	h.service.repo.FindAll(&data)

	h.Success(c, data)
}
`
}