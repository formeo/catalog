package api

type DjangoMigrationModel struct {
	ID      int    `db:"id"`
	App     string `db:"app"`
	Name    string `db:"name"`
	Applied string `db:"applied"`
}
