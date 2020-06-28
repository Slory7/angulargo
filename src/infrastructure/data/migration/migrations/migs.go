package migrations

import "github.com/slory7/angulargo/src/infrastructure/data/migration"

var MigrationVersions = []*migration.Migration{
	v201810091551(),
	v201810311029(),
}
