package observerfile

import (
	"sync"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
)

type obsFile struct {
	name     string
	filePath string
}

var obs *obsFile
var once sync.Once

// Конструктор
func NewObserverFile(obsID, filePath string) observer.ActionsObservers {
	once.Do(func() {
		obs = &obsFile{
			name:     obsID,
			filePath: filePath,
		}
	})

	return obs
}
