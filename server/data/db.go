package data

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"upper.io/bond"
	"upper.io/db/postgresql"
)

//Database structire for all connection to db
type Database struct {
	bond.Session

	Sqlx *sqlx.DB

	App     AppStore
	Release ReleaseStore
}

var (
	//DB a global refrence to db struct
	DB *Database
)

//BuildDbURL creates a URL postgress URL as string
func BuildDbURL(
	username string,
	password string,
	hosts []string,
	database string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s",
		username, password, strings.Join(hosts, ","), database)
}

//NewDB creates a db session
func NewDB(dbURL string) (*Database, error) {
	connURL, err := postgresql.ParseURL(dbURL)

	db := &Database{}
	db.Session, err = bond.Open(postgresql.Adapter, connURL)

	if err != nil {
		return nil, err
	}

	db.Sqlx = db.Session.Driver().(*sqlx.DB)

	db.App = AppStore{Store: db.Store(`apps`)}
	db.Release = ReleaseStore{Store: db.Store(`releases`)}

	DB = db

	return db, nil
}
