#!/bin/sh

sed -i "s/___REACT_APP_SOCKET_HOST___/$REACT_APP_SOCKET_HOST/g"   /usr/share/nginx/html/index.html
sed -i "s/___REACT_APP_SOCKET_DEBUG___/$REACT_APP_SOCKET_DEBUG/g" /usr/share/nginx/html/index.html

nginx -g 'daemon off;'
