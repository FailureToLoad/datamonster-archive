FROM postgres:15

# Copy initialization scripts
COPY init.sql /docker-entrypoint-initdb.d/

# Set environment variables
ENV POSTGRES_DB=records
ENV POSTGRES_USER=app
ENV POSTGRES_PASSWORD=Password1

# Expose default PostgreSQL port
EXPOSE 5432