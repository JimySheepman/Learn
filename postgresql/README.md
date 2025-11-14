# PostgreSQL

- [Follow Tutorial](https://www.postgresqltutorial.com/)

```Bash
# docker container
docker run --name postgres -e POSTGRES_PASSWORD=123456 -d -p 5432:5432 postgres
# restore data
cat dump.sql | docker exec -i postgres psql -U postgres
```
