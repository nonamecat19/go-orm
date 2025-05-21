#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <base_templ_path> <css_file_path> <js_file_path>"
    exit 1
fi

BASE_TEMPL_PATH=$1
CSS_FILE_PATH=$2
JS_FILE_PATH=$3

echo "Embedding assets from $CSS_FILE_PATH and $JS_FILE_PATH into $BASE_TEMPL_PATH"

TMP_FILE=$(mktemp)

CSS_CONTENT=$(cat "$CSS_FILE_PATH")
JS_CONTENT=$(cat "$JS_FILE_PATH")

while IFS= read -r line; do
    if [[ $line == *'<style id="main-css"></style>'* ]]; then
        echo '<style id="main-css">' >> "$TMP_FILE"
        cat "$CSS_FILE_PATH" >> "$TMP_FILE"
        echo '</style>' >> "$TMP_FILE"
    elif [[ $line == *'<script id="main-js"></script>'* ]]; then
        echo '<script id="main-js">' >> "$TMP_FILE"
        cat "$JS_FILE_PATH" >> "$TMP_FILE"
        echo '</script>' >> "$TMP_FILE"
    else
        echo "$line" >> "$TMP_FILE"
    fi
done < "$BASE_TEMPL_PATH"

mv "$TMP_FILE" "$BASE_TEMPL_PATH"

echo "Assets embedded successfully in template file."