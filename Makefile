docker-build:
	sudo docker-compose down
	sudo MYSQL_HOST=172.17.0.1 MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PASSWORD=root MYSQL_DBNAME=skyshi_gethired docker-compose up --build

docker-update:
	sudo docker tag ezizull/skyshi-gethired:latest ezizull/skyshi-gethired:$(tag)
	sudo docker tag ezizull/skyshi-gethired:latest ezizull/skyshi-gethired:latest

docker-push:
	sudo docker push ezizull/skyshi-gethired:$(tag)
	sudo docker push ezizull/skyshi-gethired:latest

docker-delete:
	sudo docker rmi -f ezizull/skyshi-gethired:$(tag)

docker-run:
	sudo docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_PORT=3306 -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DBNAME=skyshi_gethired -p 3030:3030 ezizull/skyshi-gethired:$(tag)