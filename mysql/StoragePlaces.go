package mysql

import (
	"database/sql"
)

type StoragePlaces struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
