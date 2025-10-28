package gz

import (
	"compress/gzip"
	"io"
	"net/http"
)

// compressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type compressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

// Конструктор. Возвращает экземпляр writer.
//
// Параметры:
//
// w - интерфейс ответа.
func NewCompressWriter(w http.ResponseWriter) *compressWriter {

	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

// Обработка заголовков. Возвращаются заголовки.
func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

// Реализация записи. Возвращается количество записанных байт и ошибка.
//
// Параметры:
//
// p - байты для записи.
func (c *compressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

// Добавление заголовка Content-Encoding.
//
// Параметры:
//
// statusCode - код статуса.
func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера. Возвращается ошибка.
func (c *compressWriter) Close() error {
	return c.zw.Close()
}

// compressReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
// декомпрессировать получаемые от клиента данные
type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

// --------------------------------------------------

// Конструктор. Возвращает экземпляр reader и ошибку.
//
// Параметры:
//
// r - интерфейс чтения.
func NewCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

// Реализация чтения. Возвращается количество прочитанных байт и ошибка.
//
// Параметры:
//
// p - байты для  записи.
func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

// Закрытие чтения. Возвращается ошибка.
//
// Параметры:
//
// c - указатель на reader.
func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}
