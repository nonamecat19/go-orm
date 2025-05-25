#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <base_templ_go_path> <css_file_path> <js_file_path>"
    exit 1
fi

BASE_TEMPL_GO_PATH=$1
CSS_FILE_PATH=$2
JS_FILE_PATH=$3

echo "Embedding assets from $CSS_FILE_PATH and $JS_FILE_PATH into $BASE_TEMPL_GO_PATH"

CSS_CONTENT=$(cat "$CSS_FILE_PATH")
JS_CONTENT=$(cat "$JS_FILE_PATH")

# Escape special characters for sed
CSS_CONTENT_ESCAPED=$(echo "$CSS_CONTENT" | sed 's/\\/\\\\/g; s/\//\\\//g; s/&/\\&/g; s/\$/\\$/g')
JS_CONTENT_ESCAPED=$(echo "$JS_CONTENT" | sed 's/\\/\\\\/g; s/\//\\\//g; s/&/\\&/g; s/\$/\\$/g')

sed -i 's/<style id="main-css"><\/style>/<style id="main-css">'"$CSS_CONTENT_ESCAPED"'<\/style>/g' "$BASE_TEMPL_GO_PATH"

sed -i 's/<script id="main-js"><\/script>/<script id="main-js">'"$JS_CONTENT_ESCAPED"'<\/script>/g' "$BASE_TEMPL_GO_PATH"

echo "Assets embedded successfully."
