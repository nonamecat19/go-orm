module github.com/nonamecat19/go-orm/studio

go 1.23.1

require (
	github.com/a-h/templ v0.3.865
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/lib/pq v1.10.9
	github.com/nonamecat19/go-orm/app v0.0.0-20250517165903-97e4fe84d3d7
	github.com/nonamecat19/go-orm/core v0.0.1
	github.com/nonamecat19/go-orm/orm v0.0.1
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gobuffalo/logger v1.0.6 // indirect
	github.com/gobuffalo/packd v1.0.1 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/term v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace github.com/nonamecat19/go-orm/studio => ./

replace github.com/nonamecat19/go-orm/orm => ../orm

replace github.com/nonamecat19/go-orm/core => ../core
