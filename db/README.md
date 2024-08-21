# Datamonster DB

This repo uses a Makefile to build the image and run the container.  

`make run` will run the container, mapping it to `localhost:8070`.  It also sets the admin password and app password.

## Setup

Set the `PG_PASS` environment variable for the postgres admin password.  
