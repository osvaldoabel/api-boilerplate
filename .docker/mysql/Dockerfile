FROM mysql:5.7

COPY startup_database.sql /docker-entrypoint-initdb.d/

ONBUILD RUN chown -R mysql:mysql ./dbdata
