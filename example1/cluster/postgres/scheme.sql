CREATE TABLE trips (
  id     UUID PRIMARY KEY,
  driver UUID,
  rider  UUID NOT NULL,
  status TEXT NOT NULL,
  active BOOL NOT NULL DEFAULT TRUE
);

CREATE TABLE riders (
  id UUID PRIMARY KEY
);

CREATE TABLE drivers (
  id UUID PRIMARY KEY
);
