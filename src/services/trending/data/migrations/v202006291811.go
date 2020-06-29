package migrations

import (
	"github.com/slory7/angulargo/src/infrastructure/data/db"
	"github.com/slory7/angulargo/src/infrastructure/data/migration"

	"github.com/nuveo/log"
)

func v202006291811() *migration.Migration {
	ver := "v202006291811"
	var mig = &migration.Migration{
		ID:          ver,
		Description: `Change Field(Description) length to 2000`,

		Migrate: func(db *db.Database) error {
			log.Printf("migrating: %s\n", ver)
			sql := "ALTER TABLE `gitrepo` ALTER COLUMN `description` type VARCHAR(2000)"
			_, err := db.Exec(sql)
			return err
		},

		Rollback: func(db *db.Database) error {
			log.Printf("rollback: %s\n", ver)
			sql := "ALTER TABLE `gitrepo` ALTER COLUMN `description` type VARCHAR(255)"
			_, err := db.Exec(sql)
			return err
		},
	}
	return mig
}
