package data

import "time"

//Release structure for data represents in database
type Release struct {
	ID       int64     `db:"id,pk" json:"id"`
	Note     string    `db:"note" json:"note"`
	Version  Version   `db:"version" json:"version"`
	CreateAt time.Time `db:"created_at" json:"created_at"`
}

//CollectionName returns collection name in database
func (b *Release) CollectionName() string {
	return `release`
}
