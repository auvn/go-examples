#!/bin/sh

set -e

export PGUSER="$POSTGRES_USER"

"${psql[@]}" < /scheme.sql
