FROM docker.io/alpine:latest

ARG hugoVersion=0.101.0

ENV GIT_MSG=Sync
ENV GIT_TIMEOUT=300
ENV CMS_URL=http://localhost:1337

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
    adduser --disabled-password webhook


USER webhook

ENTRYPOINT strapiwebhook serve -m "$GIT_MSG" -t $GIT_TIMEOUT -s $CMS_URL /app
