FROM alpine:latest

WORKDIR /app

COPY /build/project-layout/myapp .
# COPY /.env.common .
# COPY /.env.production .
COPY app/myapp/etc/config_production.yml app/myapp/etc/config_production.yml

# Resolve the setting `Prefork` to `true` issue.
RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ./myapp