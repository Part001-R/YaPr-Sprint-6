package flags

import (
	"flag"
	"os"
)

type ConfigT struct {
	ServerAddr       string
	BaseAddrShortURL string
	LogLevel         string
	FileStoragePath  string
	AuditFile        string
	AuditURL         string
	DSNDB            string
}

func ParseFlags() ConfigT {

	var flags = ConfigT{}

	// URL
	flag.StringVar(&flags.ServerAddr, "a", ":8080", "адрес и порт сервера")
	flag.StringVar(&flags.BaseAddrShortURL, "b", "http://localhost:8080/", "базовый адрес для коротких URL")
	flag.StringVar(&flags.LogLevel, "l", "info", "уровень логирования")
	flag.StringVar(&flags.FileStoragePath, "f", "storage.json", "хранилище ссылок")
	flag.StringVar(&flags.DSNDB, "d", "", "dsn подключения к БД")
	flag.StringVar(&flags.AuditFile, "audit-file", "", "путь к файлу-приёмнику")
	flag.StringVar(&flags.AuditURL, "audit-url", "", "URL удаленного сервера-приёмника")

	flag.Parse()

	// URL
	if envValue := os.Getenv("SERVER_ADDRESS"); envValue != "" {
		flags.ServerAddr = envValue
	}
	if envValue := os.Getenv("BASE_URL"); envValue != "" {
		flags.BaseAddrShortURL = envValue
	}
	if envValue := os.Getenv("LOG_LEVEL"); envValue != "" {
		flags.LogLevel = envValue
	}
	if envValue := os.Getenv("FILE_STORAGE_PATH"); envValue != "" {
		flags.FileStoragePath = envValue
	}
	if envValue := os.Getenv("DATABASE_DSN"); envValue != "" {
		flags.DSNDB = envValue
	}
	if envValue := os.Getenv("AUDIT_FILE"); envValue != "" {
		flags.AuditFile = envValue
	}
	if envValue := os.Getenv("AUDIT_URL"); envValue != "" {
		flags.AuditURL = envValue
	}

	return flags
}
