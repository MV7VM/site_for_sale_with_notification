package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model              // Adds some metadata fields to the table
	ID            uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	PhoneNumber   string
	WhatMaterial1 string
	Value1        string
	Address       string
	Date          string
	Comment       string
}

type NoteTask struct {
	gorm.Model              // Adds some metadata fields to the table
	ID            uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	TodayDate     string
	PhoneNumber   string
	WhatMaterial1 string
	Value1        string
	Address       string
	Date          string
	Comment       string
}
