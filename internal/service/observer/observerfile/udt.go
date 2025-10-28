package observerfile

import (
	"sync"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
)

// Представление наблюдателя
type obsFile struct {
	name     string
	filePath string
}

// экземпляр наблюдателя
var obs *obsFile

// Обеспечение единоразовой инициализации
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
