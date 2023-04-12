#!/bin/sh

help() {
    echo 'Usage: fetch-list.sh MODEL_NAME ENDPOINT'
    echo 'Fetch a list of models '
    echo
    echo 'Example:'
    echo -e 'fetch-list.sh article articles'
}

list() {
	if [ ! $# == 2 ]; then
		help
		exit 1
	fi

	local model_name=$1
	local endpoint=$2

	strapiwebhook fetch list \
		-d $SITE_DIR \
		-s $STRAPI_HOST \
		-t $STRAPI_API_TOKEN  \
		-S $SINGLE_TYPES \
		-C $COLLECTION_TYPES \
		-l $DEFAULT_LOCALE \
		-m ${model_name} \
		-e ${endpoint}
}

list $*
