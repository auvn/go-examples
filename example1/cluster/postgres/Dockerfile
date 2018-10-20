FROM library/postgres
COPY scheme.sh /docker-entrypoint-initdb.d/scheme.sh

RUN chmod +x /docker-entrypoint-initdb.d/scheme.sh
EXPOSE 5432
CMD ["postgres"]
