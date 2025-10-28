package observerurl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Part001-R/YaPr-Sprint-6/internal/service/observer"
)

// Получение ID наблюдателя
func (ou obsURL) GetID() string {
	return ou.name
}

// Передача сообщения по сети
func (ou obsURL) SendMsg(msg observer.AuditEvent) error {

	msg.URL = strings.Trim(msg.URL, `\"`)

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("ошибка json.Marshal: <%w>", err)
	}

	response, err := ou.clientHTTP.Post(ou.pathURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка отправки на удаленный сервер: <%w>", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("получен неудачный код ответа: %d", response.StatusCode)
	}

	return nil
}
