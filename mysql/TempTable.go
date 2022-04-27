package mysql

import "database/sql"

type TempTable struct {
	Id        int            `json:"id"`
	Basename  string         `json:"basename"`
	Person    string         `json:"person"`
	Size      string         `json:"size"`
	FileType  string         `json:"file_type"`
	RawData   string         `json:"raw_data"`
	Location  string         `json:"location"`
	BackedUp  bool           `json:"backed_up"`
	LogTime   int            `json:"log_time"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
