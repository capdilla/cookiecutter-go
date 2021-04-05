PATH_FOLDER=/boilerplate/template
CONTAINER_NAME=cookiecutter-maker
OUTPUT_DIR=/boilerplate/tmp
FOLDER_DEST=""

build:
	FOLDER_DEST=${FOLDER_DEST} docker-compose -f docker-compose.cookiecutter.yml up -d --build
	docker exec -it ${CONTAINER_NAME} cookiecutter ${PATH_FOLDER} --output-dir=${OUTPUT_DIR}
	docker stop ${CONTAINER_NAME}
	


 