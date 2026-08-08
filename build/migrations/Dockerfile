FROM migrate/migrate:v4.15.2
RUN apk add --no-cache postgresql-client bash

COPY scripts/init-schemas.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]