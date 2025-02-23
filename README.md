# Datamonster

An application for managing Kingdom Death campaigns.

## Local development

The compose file in the project root sets up a caddy reverse proxy container and a postgres database.  
The caddy file mimics how the project behaves in production.  

### Environment Variables

```bash
export CONN="host=localhost port=5432 user=app dbname=records password=Password1 sslmode=disable"
```

### Running Locally

This portion assumes you have docker and cmake installed.

From the root of the project run `docker compose up -d`.  
Navigate to the api directory and run `make dev`.  
Navigate to the web directory and run:  
  
```bash
nvm use
npm i
npm run dev
```
  
Navigate to `http://localhost` in your browser.  

If all went well, you'll see the Datamonster landing page.  
