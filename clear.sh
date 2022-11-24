docker ps -a | grep Exited | awk '{print $1}' | xargs docker container rm
docker image ls | grep "<none>" | awk '{print $3}' | xargs docker image rm