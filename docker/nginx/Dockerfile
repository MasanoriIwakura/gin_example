FROM nginx:1.19.8-alpine

RUN apk --update --no-cache add \
    gcc \
    make \
    libc-dev \
    g++ \
    openssl-dev \
    linux-headers \
    pcre-dev \
    zlib-dev \
    libtool \
    automake \
    autoconf \
    curl

RUN cd /opt \
    && curl -L https://github.com/openresty/headers-more-nginx-module/archive/v0.33.tar.gz | tar -xz \
    && wget -O - http://nginx.org/download/nginx-1.19.8.tar.gz | tar zxfv - \
    && mv /opt/nginx-1.19.8 /opt/nginx \
    && cd /opt/nginx \
    && ./configure --with-compat --add-dynamic-module=/opt/headers-more-nginx-module-0.33 \
    && make modules

FROM nginx:1.19.8-alpine
RUN apk --update --no-cache add openssl

COPY --from=0 /opt/nginx/objs/ngx_http_headers_more_filter_module.so /usr/lib/nginx/modules
COPY ./conf/conf.d /etc/nginx/conf.d
COPY ./conf/nginx.conf /etc/nginx/nginx.conf
COPY ./entrypoint.sh /entrypoint.sh

CMD ["/bin/sh", "/entrypoint.sh"]
