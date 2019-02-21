up/app:
	go run main.go

up/mysql:
	docker-compose up --build mysql 

down:
	docker-compose down