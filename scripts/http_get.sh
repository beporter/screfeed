#!/usr/bin/env bash
# Download a couple of example HTML files to parse into the local directory
# using curl.

OUTPUT_DIR="./html-cache"
URL_LIST=(
	"http://penny-arcade.com/comic?foo=bar"
	"http://penny-arcade.com/news"
	"http://wiki.guildwars.com/wiki/Nicholas_the_traveler"
)

if [ ! -d "${OUTPUT_DIR}" ]; then
	mkdir -p "${OUTPUT_DIR}"
fi

for FETCH_URL in "${URL_LIST[@]}"; do
	OUTPUT_FILENAME="${FETCH_URL##*/}"
	OUTPUT_FILENAME="${OUTPUT_DIR}/${OUTPUT_FILENAME%%\?*}.html"

	echo "$FETCH_URL => $OUTPUT_FILENAME"
	curl --compressed --silent --output "${OUTPUT_FILENAME}" $FETCH_URL
done
