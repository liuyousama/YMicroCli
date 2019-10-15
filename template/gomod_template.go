package template

var GomodTemplate = `
module {{.}}

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	google.golang.org/grpc v1.24.0 // indirect
)
`
