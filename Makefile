
up: gin-up
	
down: gin-down docker-down

bounce: gin-down gin-up

gin-up: docker-up
	sleep 2
	go run main.go & 

gin-down:
	pkill -l -F ./GINSVR.pid
	rm GINSVR.pid 

docker-up: 
	docker-compose -f docker-compose.yml up -d

docker-down:
	docker-compose -f docker-compose.yml down

clean-db:
	docker volume rm  simple_web_svc_db

update-swagger:
	swag init