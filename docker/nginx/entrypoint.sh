#!/bin/sh

[ -z "$HOST" ] && HOST=127.0.0.1
[ -z "$PORT" ] && PORT=3000

if [ -n "$BASIC_USER" -a -n "$BASIC_PASSWORD" ]; then
  # 改行コード
  LF=$(printf '\\\012_')
  LF=${LF%_}

  BASIC_AUTH="auth_basic \"Basic Authentication\";$LF  auth_basic_user_file \/etc\/nginx\/passwd;"
  echo $BASIC_USER:$(openssl passwd -apr1 ${BASIC_PASSWORD}) > /etc/nginx/passwd
else
  BASIC_AUTH=""
fi

sed "s/%HOST%/$HOST/g" -i /etc/nginx/conf.d/default.conf
sed "s/%PORT%/$PORT/g" -i /etc/nginx/conf.d/default.conf
sed "s/%BASIC_AUTH%/$BASIC_AUTH/g" -i /etc/nginx/conf.d/default.conf

/usr/sbin/nginx
