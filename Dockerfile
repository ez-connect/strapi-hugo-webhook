FROM docker.io/library/busybox:latest

WORKDIR /app
COPY dist/strapi-webhook-linux strapi-webhook
RUN chmod +x strapi-webhook
ENTRYPOINT /app/strapi-webhook -d /app/data
