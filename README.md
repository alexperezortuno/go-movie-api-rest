# go-api-movie

### Create a image
```
$ docker build -t ${DOCKER_USERNAME}/${DOCKER_PROJECT}:${DOCKER_TAG} . | tee image-build.log
$ docker push ${DOCKER_USERNAME}/${DOCKER_PROJECT}:${DOCKER_TAG}
```

### Create an run container
```
$ docker-compose up -d
# To run with more information
$ docker-compose --verbose up
```

### Down a container
```
$ docker-compose down
```