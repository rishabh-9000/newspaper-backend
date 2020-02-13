# Newspaper Backend

This repository only contains the backend API of the Newspaper Website.
For frontend and news crawler refer to below repositories.
Frontend - https://github.com/rishabh-9000/newspaper-frontend
Crawler - https://github.com/rishabh-9000/news_crawler
Live IP - http://35.154.105.119/

## Prerequisite

1. Golang - https://golang.org/dl/
2. MongoDB - https://www.mongodb.com/download-center/community

## Project Setup

Go to your GOPATH

```bash
cd $GOPATH
```

Clone this repository in the GOPATH

```bash
git clone https://github.com/rishabh-9000/newspaper-backend.git
```

Install Dependencies

```bash
cd newspaper-backend
dep ensure
```

Building and Running the Project

```bash
go build
./newspaper-backend
```

Environment Variables Required
1. mongo - It should contain MongoURI (for local MongoDB use mongodb://127.0.0.1:27017/)
2. db - It should contain the DB name (Ex. newspaper)
3. jwt_key - It should contain the JWT key string (Ex. key_string)
