# Datamonster

An application for managing Kingdom Death campaigns.

## Local development

### Environment Variables

#### api

```bash
export KEY="your clerk secret key"
export MODE="dev"
export CONN="host=localhost port=8070 user=app dbname=records password=Password1 sslmode=disable"
export CLIENT='http://localhost:8090'
```

### web

Create an env.local file with

```bash
VITE_CLERK_PUBLISHABLE_KEY='the publishable key from your clerk instance'
```

### nginx

I'm running on pop_OS, but these instructions should be applicable to most debian based distros.

Install nginx.

```bash
sudo apt update
sudo apt install nginx
```

Edit the default configuration.

`sudo nano /etc/nginx/sites-available/default`

```bash
server {
    listen 80;  
    server_name localhost;  

    location / {
        proxy_pass http://localhost:8090; 
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

   
    location /graphql {
        # this is paranoia from an options issue I ran into, it may not be needed
        if ($request_method = OPTIONS) {
            add_header Access-Control-Allow-Origin *;
            add_header Access-Control-Allow-Methods "GET, POST, OPTIONS";
            add_header Access-Control-Allow-Headers "Authorization, Content-Type, X-CSRF-Token";
            add_header Content-Length 0;
            add_header Content-Type text/plain;
            return 204;
        }

      
        proxy_pass http://localhost:8080; 
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host $host;
        proxy_set_header X-Forwarded-Port $server_port;
        
        # Optional: Timeout settings
        proxy_connect_timeout 60s;
        proxy_read_timeout 60s;
        proxy_send_timeout 60s;

        # Optional: Buffer settings to handle large payloads
        proxy_buffers 16 4k;
        proxy_buffer_size 2k;
    }
}

```

Check the configuration.

`sudo nginx -t`

Restart nginx.

`sudo systemctl restart nginx`

### running locally

This portion assumes you have docker and cmake installed.

Navigate to the db directory and run `make run`.
Navigate to the api directory and run `make dev`.
Navigate to the web directory and run:

```bash
nvm use
npm i
npm run dev
```

Navigate to `http://localhost` in your browser.

If all went well, you'll see the Datamonster landing page.