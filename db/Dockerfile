FROM mysql:5.7.15

MAINTAINER carsten.sielck@wilken.de

ENV MYSQL_DATABASE=banken \
    MYSQL_ROOT_PASSWORD=banken

ADD banken.sql /docker-entrypoint-initdb.d

EXPOSE 3306

