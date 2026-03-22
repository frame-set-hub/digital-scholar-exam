package usecase

import "errors"

var (
	// ErrCandidateNameRequired ชื่อว่างหลัง trim
	ErrCandidateNameRequired = errors.New("candidate name required")
	// ErrDuplicateCandidateName มีผลสอบด้วยชื่อนี้แล้ว
	ErrDuplicateCandidateName = errors.New("duplicate candidate name")
)
