#!/bin/sh

help() {
    echo 'Usage: fetch-get.sh MODEL_NAME ENDPOINT ID'
    echo 'Fetch an entry by ID'
    echo
    echo 'Example:'
    echo -e 'fetch-get.sh article articles 1'
}

get() {
	if [ ! $# == 3 ]; then
		help
		exit 1
	fi

	local model_name=$1
	local endpoint=$2
	local id=$3

	strapiwebhook fetch list \
		-d $SITE_DIR \
		-s $STRAPI_HOST \
		-t $STRAPI_API_TOKEN  \
		-S $SINGLE_TYPES \
		-C $COLLECTION_TYPES \
		-l $DEFAULT_LOCALE \
		-m ${model_name} \
		-e ${endpoint} \
		--id $3
}

get $*
