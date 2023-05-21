
#
# To install golang-migrate beforehand
# https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
#

DATABASE_PATH="storage/db"

SQLITE_URL="sqlite3://${DATABASE_PATH}"

migrate -database "$SQLITE_URL" -path db/migrations up
