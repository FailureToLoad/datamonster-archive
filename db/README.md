# datamonster.records

This repo uses a Makefile to build the image and run the container.  

`make build` will build the image.  
`make run` will run the container, mapping it to `localhost:8070`.  It also sets the admin password and app password.

The [init-db](https://github.com/FailureToLoad/datamonster.records/blob/34910d6a896477cfcd15ab6dfa01145183c1feff/init-db.sh) script sets up all the tables and users.  

## Setup

Set the `SA_PASS` environment variable for the postgres admin password.  
Set the `APP_PASS` environment variable for the app user password.  

## Related Projects

 [datamonster.api](https://github.com/FailureToLoad/datamonster.api) - The api that datamonster.web communicates against  
 [datamonster.web](https://github.com/FailureToLoad/datamonster.web) - A NextJS site which utilizes datamonster.api for critical logic and data persistence  
