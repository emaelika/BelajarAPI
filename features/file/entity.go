package file

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID        uint
	NoInvoice string
	JobID     uint
	Amount    int
	Status    string
	Token     string
	Url       string
	Timestamp time.Time `json:"created_at"`
}

type TransactionList struct {
	ID        uint
	NoInvoice string
	JobID     uint
	Status    string
	Token     string
	Url       string
	Timestamp time.Time `json:"timestamp"`
}

type Handler interface {
	Upload() echo.HandlerFunc
	Get() echo.HandlerFunc
}

type Service interface {
	Upload(file multipart.FileHeader) (string, error)
	Get(path string) (multipart.File, error)
}
