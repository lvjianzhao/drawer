#!/bin/bash

set -eu

eval "cat <<EOF
$(< /nginx-default.conf)
EOF" > /etc/nginx/conf.d/drawer.conf

exec nginx -g 'daemon off;'
