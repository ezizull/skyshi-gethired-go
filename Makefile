docker-compose:
	sudo docker-compose down 
	sudo MYSQL_HOST=172.17.0.1 MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PASSWORD=root MYSQL_DBNAME=skyshi_gethired docker-compose up

docker-update:
	sudo docker tag gethiredgo_restapi:latest ezizull/skyshi-gethired:$(tag)
	sudo docker tag gethiredgo_restapi:latest ezizull/skyshi-gethired:latest

docker-push:
	sudo docker push ezizull/skyshi-gethired:$(tag)
	sudo docker push ezizull/skyshi-gethired:latest

docker-run:
	sudo docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DBNAME=skyshi_gethired -p 8090:3030 ezizull/skyshi-gethired:$(tag)