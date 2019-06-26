FROM nginx:stable-alpine
LABEL version="1.0"

COPY web/nginx.conf /etc/nginx/nginx.conf

WORKDIR /usr/share/nginx/html
COPY web/dist/frontend/ .