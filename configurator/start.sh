#!/bin/bash

set -e

PRE_START=/app/prestart.bash
if [ -f $PRE_START ] ; then
  source $PRE_PRE_START
fi

# Start Supervisor with nginx and uWSGI
exec /usr/bin/supervisord -c /etc/supervisor/conf.d/supervisord.conf
