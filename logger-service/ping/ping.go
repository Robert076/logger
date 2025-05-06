package ping

import "time"

type Ping struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
