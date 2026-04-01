package database

import (
	"log"

	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Endpoint struct {
	EndpointID string
	Value      string
	Method     string
	Type       string
	Bypass     int
}

func SyncEndpoints(db *gorm.DB, r *gin.Engine) {

	routes := r.Routes()

	for _, route := range routes {

		var count int64

		db.Table("endpoint").
			Where("value = ? AND method = ?", route.Path, route.Method).
			Count(&count)

		if count == 0 {

			err := db.Table("endpoint").Create(map[string]interface{}{
				"endpoint_id": idgen.NewEndpointID(),
				"value":       route.Path,
				"method":      route.Method,
				"type":        "API",
				"bypass":      0,
			}).Error

			if err != nil {
				log.Println("endpoint insert error:", err)
			} else {
				log.Println("endpoint registered:", route.Method, route.Path)
			}
		}
	}
}