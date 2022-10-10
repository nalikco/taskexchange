# Running migrations:
Up:\
```:~$ migrate -path ./schema -database "postgres://user:pass@host:port/dbname?sslmode=disable" up```\
Down:\
```:~$ migrate -path ./schema -database "postgres://user:pass@host:port/dbname?sslmode=disable" up```

# Settings:
Server: **.env** file\
Client: **client/.env** file

# Building:
Server:\
```:~$ go build cmd/main.go```\
Client:\
```:~$ cd client && npm run build```

# Running:
```:~$ ./main```