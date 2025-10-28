package observerurl

import (
	"net/http"
	"sync"
	"time"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
)

// Представление наблюдателя
type obsURL struct {
	name       string
	pathURL    string
	clientHTTP *http.Client
}

// Экземпляр наблюдателя
var obs *obsURL

// Обеспечение единоразовой инициализации
var once sync.Once

// Конструктор. Возвращается интерфейс.
//
// Параметры:
//
// obsID - ID наблюдателя.
// obsPath - URL наблюдателя.
func NewObserverURL(obsID, obsPath string) observer.ActionsObservers {
	once.Do(func() {

		client := &http.Client{
			Timeout: 2 * time.Second,
		}

		obs = &obsURL{
			name:       obsID,
			pathURL:    obsPath,
			clientHTTP: client,
		}
	})

	return obs
}
