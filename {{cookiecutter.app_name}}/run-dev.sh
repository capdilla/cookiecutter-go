#!/bin/bash
set -e 

# Always rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

PROFILE=$1
if [ -z "$1" ]; then
  echo "an AWS profile is required to continue"
  exit
fi
shift
echo "-------"
echo "starting {{cookiecutter.app_name}}, you can pass flags to docker compose after the profile argument normally"
echo "-------"
export AWS_PROFILE=${PROFILE}
export AWS_ACCESS_KEY_ID=$(aws configure get ${AWS_PROFILE}.aws_access_key_id)
export AWS_SECRET_ACCESS_KEY=$(aws configure get ${AWS_PROFILE}.aws_secret_access_key)
export FCM_CREDENTIALS=$(cat credentials.json)
echo "[+] Run containers ${@} "
docker-compose up ${@}

echo "[+] Cleaning up stopped containers..."
docker ps --all --filter status=exited -q | xargs docker rm -v;