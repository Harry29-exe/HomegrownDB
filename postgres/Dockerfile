FROM postgres:13.2

#COPY initdb /docker-entrypoint-initdb.d
EXPOSE 5432
ENV POSTGRES_PASSWORD=123

CMD ["postgres", "-c", "log_statement=all"]