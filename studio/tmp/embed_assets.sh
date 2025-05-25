#!/bin/bash
CSS_CONTENT=$(cat tmp/main.css.content)
JS_CONTENT=$(cat tmp/main.js.content)
sed -i "s|<style id=\"main-css\"></style>|<style id=\"main-css\">${CSS_CONTENT}</style>|g" internal/view/layout/base_templ.go
sed -i "s|<script id=\"main-js\"></script>|<script id=\"main-js\">${JS_CONTENT}</script>|g" internal/view/layout/base_templ.go
