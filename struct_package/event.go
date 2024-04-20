package struct_package

type Event struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Photos      []EventPhoto `json:"photos"`
	Video       *string      `json:"video"`
	Map         *string      `json:"map"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type EventPhoto struct {
	ID        int64  `json:"id"`
	EventID   int64  `json:"event_id"`
	File      []byte `json:"file"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
