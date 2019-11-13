To initialize go mod for our app
- go mod init github.com/ijogleka/dockerized-http-server

To test our app locally
- go run main.go

To build our docker container image
- docker build -t . go-docker 
  where go-docker is our image name
  
To run our image in (as) a container
- docker run -d -p 8080:8080 go-docker 
  -d is for running this in the background
  -p 8080:8080 is for mapping the port 8080 on the container to the localhost 8080
  
To test our server out
- Open up a browser and type in `localhost:8080`
- You should see the output 
  `Hello world from the dockerized http server`
  
  
Other Notes
- To look up all running docker containers 
  docker container ls
- To stop a container
  docker container stop <container_id>
- To lookup all images
  docker images

