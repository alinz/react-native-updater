package data

import "time"

//App struct for storing the basic information about each app
type App struct {
	ID       SecureID  `db:"id,pk" json:"id"`
	Name     string    `db:"name" json:"name"`
	CreateAt time.Time `db:"cretad_at" json:"created_at"`
}

//CollectionName returns collection name in database
func (b *App) CollectionName() string {
	return `apps`
}
