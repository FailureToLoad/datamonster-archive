run:
	docker run \
		-e POSTGRES_PASSWORD=${PG_PASS} \
		-e POSTGRES_USER='app' \
		-e POSTGRES_DB='records' \
		-p 127.0.0.1:8070:5432  \
		-d \
		--name dm-db postgres
