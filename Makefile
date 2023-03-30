docker-compose:
	sudo docker-compose down 
	sudo MYSQL_HOST=172.17.0.1 MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PASSWORD=root MYSQL_DBNAME=skyshi_gethired docker-compose up
