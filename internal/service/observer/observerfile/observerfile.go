package observerfile

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/logger"
	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
	"go.uber.org/zap"
)

// Получение ID наблюдателя
func (of obsFile) GetID() string {
	return of.name
}

// Сохранение сообщения в файл
func (of obsFile) SendMsg(msg observer.AuditEvent) error {

	file, err := os.OpenFile(of.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Log.Error("Ошибка открытия файла аудита", zap.Error(err))
		return fmt.Errorf("ошибка открытия файла аудита: <%v>", err)
	}
	defer file.Close()

	msg.URL = strings.Trim(msg.URL, `\"`)

	data, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Error("Ошибка json.Marshal", zap.Error(err))
		return fmt.Errorf("оОшибка json.Marshal: <%v>", err)
	}
	file.WriteString(string(data) + "\n")
	return nil
}
