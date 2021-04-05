PATH_FOLDER=/boilerplate
CONTAINER_NAME=cookiecutter-maker

build:
	docker-compose -f docker-compose.cookiecutter.yml up -d --build
	docker exec -it ${CONTAINER_NAME} cookiecutter ${PATH_FOLDER}
	docker stop ${CONTAINER_NAME}

eject:
	rm
	


 