package migrations

import (
	"github.com/slory7/angulargo/src/infrastructure/data/db"
	"github.com/slory7/angulargo/src/infrastructure/data/migration"
	"github.com/slory7/angulargo/src/infrastructure/datamodels"

	"github.com/nuveo/log"
)

func v201810091551() *migration.Migration {
	ver := "v201810091551"
	var mig = &migration.Migration{
		ID:          ver,
		Description: `Add User(Age) field`,

		Migrate: func(db *db.Database) error {
			log.Printf("migrating: %s\n", ver)
			err := db.Sync(new(datamodels.User))
			return err
		},

		Rollback: func(db *db.Database) error {
			log.Printf("rollback: %s\n", ver)
			sql := "ALTER TABLE `user` DROP COLUMN `age` "
			_, err := db.Exec(sql)
			return err
		},
	}
	return mig
}
