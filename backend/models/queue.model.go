package models

import "time"

type StatusQ string
type PriorityQ string

const (
	StatusPending    StatusQ = "pending"
	StatusProcessing StatusQ = "processing"
	StatusCompleted  StatusQ = "completed"
	StatusFailed     StatusQ = "failed"
	StatusCancelled  StatusQ = "cancelled"
)

const (
	PriorityHigh PriorityQ = "high"
	PriorityMid  PriorityQ = "mid"
	PriorityLow  PriorityQ = "low"
)

type (
	Queue struct {
		Status       StatusQ
		Priority     PriorityQ
		Source_Url   string
		Filename     string
		Filesize     float32
		Progress     float32
		Attempts     uint8
		MaxAttempts  uint8
		ErrorMessage string
		CreatedAt    time.Time
		StartedAt    time.Time
		CompletedAt  time.Time
		Model
	}
)
