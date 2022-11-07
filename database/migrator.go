package database

import "github.com/ario-team/glassnode-api/schema"

func Migrate() {
	Connection.AutoMigrate(&schema.Endpoint{}, &schema.TopGroup{}, &schema.MiddleGroup{})
}
