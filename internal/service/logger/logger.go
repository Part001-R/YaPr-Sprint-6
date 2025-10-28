package logger

import (
	"go.uber.org/zap"
)

// Глобальный логгер.
var Log *zap.Logger = zap.NewNop()

// Инициализация логгера. Возвращается ошибка.
//
// Параметры:
//
// level - уровень логирования.
func Initialize(level string) error {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	// устанавливаем синглтон
	Log = zl
	return nil
}
