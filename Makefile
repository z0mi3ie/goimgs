up/app:
	go run main.go

up/mysql:
	docker-compose up --build mysql 

up/imgserver:
	docker-compose up --build imgserver

down:
	docker-compose down