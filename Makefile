
migrate: go get -u github.com/golang-migrate/migrate/v4

create_migrate_file : migrate create -ext sql -dir migrations -seq create_users
