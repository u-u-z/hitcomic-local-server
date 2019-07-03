package main

import "time"

// BasicModel : basic of all database models.
type BasicModel struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tickets :
type Tickets struct {
	BasicModel
	Key   string
	Type  uint
	Times uint
}

// Logs : logs for ticket
type Logs struct {
	BasicModel
	Key    string
	Result uint
}

// CertPicture : logs for ticket
type CertPicture struct {
	BasicModel
	Tickets  Tickets `gorm:"ForeignKey:TicketID"`
	TicketID uint64
	Path     string `sql:"type:VARCHAR(620) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"`
}
