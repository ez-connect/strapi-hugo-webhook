FROM docker.io/nginx:alpine-slim

ARG hugoVersion=0.110.0

ENV STRAPI_HOST=http://localhost:1337
ENV STRAPI_API_TOKEN=api-token
ENV SITE_DIR=/home/web
ENV TEMPLATE_DIR=/home/web/.config/template
ENV DEFAULT_LOCALE=en
ENV SINGLE_TYPES=site,home,nav,about
ENV COLLECTION_TYPES=section,contributor,article,document,career,project,page,resume
ENV CMD="echo 'build the site'; hugo --gc --minify;"
ENV DEBOUNCED_CMD="echo 'debounced cmd'"
ENV DEBOUNCED_TIMEOUT=300

COPY --chmod=001 build/linux/strapiwebhook /usr/local/bin/
COPY --chmod=001 data/oci/bin/* /usr/local/bin/
COPY data/oci/config/nginx/ /etc/nginx/conf.d/

WORKDIR /home

RUN set -eux; \
	# Install git
	apk add --no-cache make curl git openssh; \
	# Passing SSH option to git: https://stackoverflow.com/a/38474400
	git config --global core.sshCommand 'ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no'; \
	# Install hugo
	apk add --no-cache libc6-compat libstdc++; \
	curl -Lo /tmp/hugo.tar.gz https://github.com/gohugoio/hugo/releases/download/v${hugoVersion}/hugo_extended_${hugoVersion}_Linux-64bit.tar.gz; \
	tar -C /tmp -xf /tmp/hugo.tar.gz; \
	mv /tmp/hugo /usr/local/bin; \
	rm -rf /tmp/*; \
	hugo version

EXPOSE 80 8080

CMD nginx -g 'daemon on;'; \
	strapiwebhook serve \
		-d=$SITE_DIR \
		-T=$TEMPLATE_DIR \
		-C=$COLLECTION_TYPES \
		-S=$SINGLE_TYPES \
		-l=$DEFAULT_LOCALE \
		-s=$STRAPI_HOST \
		-t=$TRAPI_API_TOKEN \
		-c="$CMD" \
		--debounced-cmd="$DEBOUNCED_CMD" \
		--debounced-timeout=$DEBOUNCED_TIMEOUT
