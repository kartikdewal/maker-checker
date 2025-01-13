package common

import "errors"

var (
	ErrProfileNotFound         = errors.New("profile not found")
	ErrDocumentNotFound        = errors.New("document not found")
	ErrDocumentRequestNotFound = errors.New("document request not found")
)
