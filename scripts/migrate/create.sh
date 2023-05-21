
#
# To install golang-migrate beforehand
# https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
#

FILE_NAME=$1
migrate create -ext sql -dir db/migrations -seq "$FILE_NAME"
