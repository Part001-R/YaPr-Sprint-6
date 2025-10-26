package observerurl

import (
	"net/http"
	"sync"
	"time"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
)

type ObsURL struct {
	name       string
	pathURL    string
	clientHTTP *http.Client
}

var obs *ObsURL
var once sync.Once

// Конструктор
func NewObserverURL(obsID, obsURL string) observer.ActionsObservers {
	once.Do(func() {

		client := &http.Client{
			Timeout: 2 * time.Second,
		}

		obs = &ObsURL{
			name:       obsID,
			pathURL:    obsURL,
			clientHTTP: client,
		}
	})

	return obs
}
