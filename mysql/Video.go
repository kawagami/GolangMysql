package mysql

import (
	"database/sql"
)

type Video struct {
	Id        int            `json:"id"`
	ActressId int            `json:"actress_id"`
	PlaceId   int            `json:"place_id"`
	Name      string         `json:"name"`
	Size      string         `json:"size"`
	Location  string         `json:"location"`
	BackedUp  string         `json:"backed_up"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
