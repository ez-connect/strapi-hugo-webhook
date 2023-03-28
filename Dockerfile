FROM docker.io/nginx:alpine-slim

ARG hugoVersion=0.110.0

ENV STRAPI_HOST='http://localhost:1337'
ENV STRAPI_API_TOKEN='api-token'
ENV HUGO_SITE_DIR='web'
ENV DEFAULT_LOCALE='en'
ENV SINGLE_TYPES='site,home,nav,about'
ENV COLLECTION_TYPES='contributor,article,document,career,project,page,resume'
ENV GIT_COMMIT='Sync CMS'
ENV GIT_TIMEOUT='300'

COPY --chmod=001 build/linux/strapiwebhook /usr/local/bin/

RUN set -eux; \
    # Install git
    apk add --no-cache make curl git; \
    # Passing SSH option to git: https://stackoverflow.com/a/38474400
    git config --global core.sshCommand 'ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no'; \
    # Install hugo
	apk add --no-cache libc6-compat libstdc++; \
    curl -Lo /tmp/hugo.tar.gz https://github.com/gohugoio/hugo/releases/download/v${hugoVersion}/hugo_extended_${hugoVersion}_Linux-64bit.tar.gz; \
	tar -C /tmp -xf /tmp/hugo.tar.gz; \
    mv /tmp/hugo /usr/local/bin; \
	rm -rf /tmp/*; \
    hugo version; \
	# Nginx config
	echo -e '\n\
		server {\n\
			listen 80 default_server;\n\
			listen [::]:80 default_server;\n\
			gzip on;\n\
			root /home/webhook/web/public;\n\
		}\
	' > /etc/nginx/conf.d/default.conf;

WORKDIR /home/webhook

CMD nginx -g 'daemon on;'; \
	strapiwebhook serve \
		--strapi=$STRAPI_HOST \
		--token=$STRAPI_API_TOKEN \
		--dir=$HUGO_SITE_DIR \
		--locale=$DEFAULT_LOCALE \
		--single=$SINGLE_TYPES \
		--collection=$COLLECTION_TYPES \
		--commit="$GIT_MSG" \
		--timeout=$GIT_TIMEOUT
