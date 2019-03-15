local/run/app:
	go run main.go

local/test/unittest:
	go test ./...

up/goimgs:
	docker-compose up --build app

rebuild/goimgs:
	docker-compose down app
	docker-compose up --build app

up/mysql:
	docker-compose up --build mysql 

up/imgserver:
	docker-compose up --build imgserver

down:
	docker-compose down