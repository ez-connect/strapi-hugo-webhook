FROM docker.io/alpine:latest

ARG hugoVersion=0.101.0

ENV STRAPI_HOST='http://localhost:1337'
ENV STRAPI_API_TOKEN='api-token'
ENV HUGO_SITE_DIR='/home/webhook'
ENV DEFAULT_LOCALE='en'
# ENV SINGLE_TYPES=
# ENV COLLECTION_TYPES=
ENV GIT_COMMIT='Sync CMS'
ENV GIT_TIMEOUT='300'

COPY dist/linux/strapiwebhook .

RUN chmod +x strapiwebhook && \
    mv strapiwebhook /usr/bin/ && \
    # Install git
    apk add --no-cache make git less openssh && \
    # Passing SSH option to git: https://stackoverflow.com/a/38474400
    git config --global core.sshCommand 'ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' && \
    # Install nginx
    apk add --no-cache nginx && \
    # Install hugo
    apk add --no-cache libc6-compat libstdc++ && \
    wget https://github.com/gohugoio/hugo/releases/download/v${hugoVersion}/hugo_extended_${hugoVersion}_Linux-64bit.tar.gz && \
    tar -xf hugo_extended_${hugoVersion}_Linux-64bit.tar.gz && \
    rm LICENSE README.md *.gz && \
    chmod +x hugo && mv hugo /usr/bin && \
    hugo version && \
    # User
    adduser -h home/webhook --disabled-password webhook

WORKDIR /home/webhook

USER webhook

ENTRYPOINT strapiwebhook serve \
	--strapi $STRAPI_HOST \
	--token $STRAPI_API_TOKEN \
	--token $STRAPI_API_TOKEN \
	--dir $HUGO_SITE_DIR \
	--locale $DEFAULT_LOCALE \
	--commit "$GIT_MSG" \
	--timeout $GIT_TIMEOUT
