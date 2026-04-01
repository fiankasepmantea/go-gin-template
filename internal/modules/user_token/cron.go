package usertoken

import (
	"log"
	"time"
)

func StartCleanupJob(repo *Repository) {

	ticker := time.NewTicker(1 * time.Hour)

	go func() {
		for range ticker.C {
			err := repo.DeleteExpired()
			if err != nil {
				log.Println("cleanup failed:", err)
			} else {
				log.Println("expired tokens cleaned")
			}
		}
	}()
}