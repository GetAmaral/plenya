package workers

import (
	"log"
	"time"

	"github.com/plenya/api/internal/services"
)

type NotificationWorker struct {
	service  *services.NotificationService
	interval time.Duration
	stopChan chan bool
}

func NewNotificationWorker(service *services.NotificationService, interval time.Duration) *NotificationWorker {
	return &NotificationWorker{
		service:  service,
		interval: interval,
		stopChan: make(chan bool),
	}
}

// Start inicia o worker em background
func (w *NotificationWorker) Start() {
	log.Println("🔔 Notification worker started")
	log.Printf("   Check interval: %v", w.interval)

	// Run immediately on start
	w.checkSubscriptions()

	// Then run periodically
	ticker := time.NewTicker(w.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.checkSubscriptions()
			case <-w.stopChan:
				ticker.Stop()
				log.Println("🔔 Notification worker stopped")
				return
			}
		}
	}()
}

// Stop para o worker
func (w *NotificationWorker) Stop() {
	w.stopChan <- true
}

// checkSubscriptions verifica subscriptions e cria notificações
func (w *NotificationWorker) checkSubscriptions() {
	log.Println("🔔 Checking subscriptions for notifications...")

	if err := w.service.CheckAndCreateSubscriptionNotifications(); err != nil {
		log.Printf("❌ Error checking subscriptions: %v", err)
		return
	}

	log.Println("✅ Subscription notifications check completed")
}
