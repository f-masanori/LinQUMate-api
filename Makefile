DB_SERVICE:=db
DBNAME:=linqumate
DOCKER_DNS:=db
DBPASS:=password
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=$(DBPASS)

dcu:
	docker-compose up
# データベース
mysql/client:
	docker-compose exec $(DB_SERVICE) mysql -uroot -hlocalhost -p$(DBPASS) $(DBNAME)

mysql/init:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -p$(DBPASS) \
		-e "create database \`$(DBNAME)\`"

__mysql/drop:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -p$(DBPASS) \
		-e "drop database \`$(DBNAME)\`"

# マイグレーション
MIGRATION_SERVICE:=migration
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline