```
export POSTGRESQL_URL='postgres://localhost:5432/payments?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```