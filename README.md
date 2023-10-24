
# Person Predictor

REST-API server with Postgres database for store and manage Persons.

Create Person by Name, Surname and Patronymic. Age, Gender and Country will be predicted from extrnal source. Data will be added to database.

Get and Update Person by any field.

Delete Person by Id.

## Run Locally

Clone the project:

```bash
  git clone https://github.com/Rustam2202/person-predictor.git
```

Config parameters in ```app.env``` and start server:

```bash
  make run
```

## Tech Stack

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org) with [GORM](https://gorm.io/)
- [Zap Logger](https://github.com/uber-go/zap)
- [Swagger](https://github.com/swaggo/swag)
- Graceful server shutdown

## Roadmap
- Mock tests for repository
- Refactor get Age, Gender and Country funcs in server
