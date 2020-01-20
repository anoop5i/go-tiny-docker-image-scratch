Example to create a tiny docker image for GO microservices with scratch base image

Steps:

    Build the exeutable for the environment
    env GOOS=linux GOARCH=386 go build main.go //for linux

    Build the docker image
    docker build my-go-scratch .  //creates an image with size ~6MB

    Run the image
    docker run -it -p8081:8081 my-go-scratch 
