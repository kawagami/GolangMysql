package mysql

import (
	"database/sql"
)

type Comic struct {
	Id        string         `json:"id"`
	AuthorId  string         `json:"author_id"`
	PlaceId   string         `json:"place_id"`
	Name      string         `json:"name"`
	Size      string         `json:"size"`
	Location  string         `json:"location"`
	BackedUp  string         `json:"backed_up"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
