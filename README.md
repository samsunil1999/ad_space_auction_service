# ad_space_auction_service

In order to run this application:
- create mySQL DB with DSN credentials mentioned below.
- run commands in your app path
```
docker build -t seller-app-img:latest -f Dockerfile .
docker run -it -p 127.0.0.1:8080:8080 --name seller-app seller-app-img:latest
```

DSN credentials:
- host: `localhost`
- port: `3306`
- database: `sellerapp`
- password: `password`

**Note:** I've also shared the postman collection for the same 