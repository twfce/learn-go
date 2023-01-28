
```bash
# Run MariaDB in docker container for testing
docker run --detach --name golang-access-mariadb --env MARIADB_DATABASE=test --env MARIADB_USER=test --env MARIADB_PASSWORD=test --env MARIADB_ROOT_PASSWORD=test -p 3306:3306  mariadb:latest
```

```bash
# Execute SQL script to create tables
mysql --host=localhost --user=test --password=test --port 3306 --protocol=TCP < db.sql
```