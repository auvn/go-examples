#!/bin/sh

set -e

export PGUSER="$POSTGRES_USER"

"${psql[@]}" <<- 'EOSQL'
CREATE TABLE trips (
  id     UUID PRIMARY KEY,
  driver UUID NOT NULL,
  rider  UUID NOT NULL,
  status TEXT NOT NULL
);

CREATE TABLE riders (
  id UUID PRIMARY KEY
);

CREATE TABLE drivers (
  id UUID PRIMARY KEY
);
EOSQL
