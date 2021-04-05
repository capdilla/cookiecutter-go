# {{cookiecutter.app_name}}

## Usage example

`curl http://localhost:{{cookiecutter.app_dev_port}}/health`

## Development setup

### Credentials

#### set your aws credentials in your environment, or you can use autoenv as:

```.env
  AWS_ACCESS_KEY_ID: YOUR_AWS_ACCESS_KEY_ID
  AWS_SECRET_ACCESS_KEY: YOUR_AWS_SECRET_ACCESS_KEY
  AWS_DEFAULT_REGION: us-east-2
```

Running Docker Image:
` docker-compose up`

### IF you are coming from an old version run:

`docker-compose down && docker-compose up --force-recreate --remove-orphans`

### Ports

#### Exposed ports are:

`{{cookiecutter.app_dev_port}}` api interface
`{{cookiecutter.db_dev_port}}` postgres

## Migrations

### Run migrations :

### install:

```
go get -u github.com/pressly/goose/cmd/goose
```

### create migration:

```bash
make migration migration-name
```

### run migrate-up:

```bash
make migrate-up 
```

for more info about migration see `Makefile`
