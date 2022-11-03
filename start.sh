#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOUECE" -verbose up

echo "start app"
exec "$@"