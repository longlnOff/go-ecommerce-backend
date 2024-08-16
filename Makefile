run:
	go run cmd/server/main.go

mysql_docker:
	docker run -it --name mysqlgo8 -e MYSQL_DATABASE=shopdevgo -e MYSQL_ROOT_PASSWORD=root1234 -p 3306:3306 -d mysql:latest

start_mysql:
	docker start mysqlgo8

stop_mysql:
	docker stop mysqlgo8

redis_docker:
	docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest

start_redis:
	docker start redis-stack-server

stop_redis:
	docker stop redis-stack-server

.PHONY: run mysql_docker start_mysql stop_mysql redis_docker start_redis stop_redis