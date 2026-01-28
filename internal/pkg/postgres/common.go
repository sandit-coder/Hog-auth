package postgres

import (
	"fmt"
)

type Config struct {
	Dsn string
}

func GetPostgresConnectionString(dsn string) string {
	return fmt.Sprintf(dsn)
}
