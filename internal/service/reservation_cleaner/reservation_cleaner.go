package reservation_cleaner

import (
	"context"
	"log"
	"time"

	"product_service/internal/repository"
)

// ReservationCleaner отвечает за очистку истекших резервов
type ReservationCleaner struct {
	repo repository.ProductRepository
}

// NewReservationCleaner создаёт новый сервис очистки
func NewReservationCleaner(repo repository.ProductRepository) *ReservationCleaner {
	return &ReservationCleaner{repo: repo}
}

// Start запускает процесс очистки в фоновом режиме
func (rc *ReservationCleaner) Start() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Cleaning expired reservations...")
		err := rc.repo.ClearExpiredReservations(context.Background())
		if err != nil {
			log.Printf("Error clearing reservations: %v", err)
		} else {
			log.Println("Expired reservations cleaned successfully")
		}
	}
}
