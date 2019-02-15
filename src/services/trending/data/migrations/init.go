package migrations

import (
	"services/infrastructure/data"
	"services/infrastructure/data/migration"
	m "services/trending/datamodels"

	"github.com/nuveo/log"
)

func InitMigration() *migration.Migration {
	ver := "v2019901311559"
	var mig = &migration.Migration{
		ID:          ver,
		Description: `Init trending everything`,

		Migrate: func(db *data.Database) error {
			log.Printf("init migrating: %s\n", ver)

			if err := db.Sync(new(m.GitRepo)); err != nil {
				return err
			}
			if err := db.Sync(new(m.GitRepoTrending)); err != nil {
				return err
			}
			sql := `ALTER TABLE gitrepo
			ADD CONSTRAINT fk_gitrepo
			FOREIGN KEY (git_trending_id)
			REFERENCES gitrepo_trending(id)
			ON DELETE CASCADE;`
			if _, err := db.Exec(sql); err != nil {
				return err
			}

			return nil
		},

		Rollback: func(db *data.Database) error {
			log.Printf("init rollback: %s\n", ver)

			if err := db.DropTable(new(m.GitRepo)); err != nil {
				return err
			}
			if err := db.DropTable(new(m.GitRepoTrending)); err != nil {
				return err
			}
			return nil
		},
	}
	return mig
}
