package observer

// Общий интерфейс для наблюдателей
type ActionsObservers interface {
	SendMsg(msg AuditEvent) error
	GetID() string
}
