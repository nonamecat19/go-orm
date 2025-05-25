BINPATH = ./bin/myapp

.PHONY: build
build: build-css build-js embed-assets-in-template build-templ build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/main.go

.PHONY: build-templ
build-templ:
	templ generate

.PHONY: build-css
build-css:
	npm --prefix web run build:css -- --minify

.PHONY: build-js
build-js:
	npm --prefix web run build:js -- --minify

.PHONY: watch
watch:
	$(MAKE) -j5 watch-app watch-templ watch-css watch-js watch-assets

.PHONY: watch-app
watch-app:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) build-app" \
	--build.bin "$(BINPATH)" \
	--build.include_ext "go" \
	--build.exclude_dir "bin,web"

.PHONY: watch-templ
watch-templ:
	templ generate \
	--watch \
	--proxy "http://localhost:8080" \
	--open-browser=false

.PHONY: watch-css
watch-css:
	npm --prefix web run build:css -- --watch=always

.PHONY: watch-js
watch-js:
	npm --prefix web run build:js -- --watch=forever

.PHONY: watch-assets
watch-assets:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) embed-assets-in-template && templ generate --notify-proxy" \
	--build.bin "true" \
	--build.exclude_dir "" \
	--build.include_dir "web/public/assets" \
	--build.include_ext "css,js" \
	--build.delay "100"

.PHONY: fmt
fmt:
	templ fmt internal/view

.PHONY: embed-assets-in-template
embed-assets-in-template:
	@echo "Embedding assets into base.templ..."
	@chmod +x scripts/embed_assets_in_template.sh
	@./scripts/embed_assets_in_template.sh internal/view/layout/base.templ web/public/assets/main.css web/public/assets/main.js

.PHONY: embed-assets
embed-assets:
	@echo "Embedding assets into base_templ.go..."
	@chmod +x scripts/embed_assets.sh
	@./scripts/embed_assets.sh internal/view/layout/base_templ.go web/public/assets/main.css web/public/assets/main.js

.PHONY: install-deps
install-deps:
	npm --prefix web install
	go mod download
