FROM docker.io/alpine:latest

ARG hugoVersion=0.94.0

WORKDIR /app

COPY dist/strapi-webhook-linux strapi-webhook

RUN chmod +x strapi-webhook && \
    # Install git
    apk add --update-cache --no-cache make git less openssh && \
    # Passing SSH option to git: https://stackoverflow.com/a/38474400
    git config --global core.sshCommand 'ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' && \
    # # Setup SSH
    # if [ -f id_rsa ]; then \
    #     mv id_rsa* ~/.ssh/; \
    # else \
    #     ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa; \
    # fi && \
    # Install hugo
    apk add --no-cache libc6-compat libstdc++ && \
    wget https://github.com/gohugoio/hugo/releases/download/v${hugoVersion}/hugo_extended_${hugoVersion}_Linux-64bit.tar.gz && \
    tar -xf hugo_extended_${hugoVersion}_Linux-64bit.tar.gz && \
    rm LICENSE README.md *.gz && \
    chmod +x hugo && mv hugo /usr/bin && \
    hugo version

# ENTRYPOINT /app/strapi-webhook -d /app/data
