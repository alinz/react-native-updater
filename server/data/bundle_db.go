package data

import "time"

//Bundle structure for each bundle represents in database
type Bundle struct {
	ID        SecureID  `db:"id,pk" json:"id"`
	ReleaseID int64     `db:"release_id" json:"release_id"`
	Hash      string    `db:"hash" json:"hash"`
	Name      string    `db:"name" json:"name"`
	Type      Type      `db:"type" json:"type"`
	CreatedAt time.Time `db:"cretad_at" json:"created_at"`
}

//CollectionName returns collection name in database
func (b *Bundle) CollectionName() string {
	return `bundles`
}
