package services

import (
	"21-api/features/file"
	"errors"
	"mime/multipart"
)

type TransactionService struct {
}

func New() file.Service {
	return &TransactionService{}
}

func (at *TransactionService) Upload(file multipart.FileHeader) (string, error) {

	return "result", errors.New("")
}

func (ct *TransactionService) Get(transactionID string) (multipart.File, error) {

	return nil, errors.New("")
}
