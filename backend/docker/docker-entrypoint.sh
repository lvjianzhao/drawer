#!/bin/bash

set -eu

eval "cat <<EOF
$(< /tmp/config_tpl.yaml)
EOF" > /opt/drawer/config.yaml

exec /opt/drawer/drawer server