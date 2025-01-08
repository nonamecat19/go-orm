migration-create:
	go run ./cli/main.go migration:create --config ./app/db.yaml

app-run:
	go run ./app/main.go