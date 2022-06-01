package datastruct

import "time"

type Audit struct {
	CreatedAt   time.Time `db:"created_at"`
	CreatedFrom string    `db:"created_from"`
}
