FROM mysql:5.7
# WORKDIR /home/workspace

COPY startup_database.sql /docker-entrypoint-initdb.d/

# ONBUILD RUN chown -R mysql:mysql /var/lib/pgsql/data%