FROM mysql:5.6

RUN { \
  echo '[mysqld]'; \
  echo 'character-set-server = utf8'; \
  echo 'skip-character-set-client-handshake'; \

  echo '[client]'; \
  echo 'default-character-set = utf8'; \

  echo '[mysql]'; \
  echo 'no-auto-rehash'; \
  echo 'default-character-set = utf8'; \

  echo '[mysqldump]'; \
  echo 'default-character-set = utf8'; \

} > /etc/mysql/conf.d/my.cnf
