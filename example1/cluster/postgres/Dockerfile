FROM library/postgres
COPY scheme.sql /scheme.sql
COPY scheme.sh /docker-entrypoint-initdb.d/scheme.sh

RUN chmod +x /docker-entrypoint-initdb.d/scheme.sh
