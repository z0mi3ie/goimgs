up/app:
	go run main.go

up/goimgs:
	#docker-compose up --build app
	docker-compose up app

rebuild/goimgs:
	docker-compose down app
	docker-compose up --build app

up/mysql:
	docker-compose up --build mysql 

up/imgserver:
	docker-compose up --build imgserver

down:
	docker-compose down