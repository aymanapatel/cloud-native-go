# Base image that Docker image is based on. 
# Lightweight Alpine Linux distro
FROM golang:1.8.1-alpine


# Since alpine is barebones distro
# Therefore git and bash is to be installed by explicitly

RUN apk update && apk upgrade && apk add --no-cache bash git 

# Install Gin framework
RUN go get github.com/gin-gonic/gin

# Copy image's filesystem or working directory to current directory
ENV SOURCES /home/ayman/Desktop/GOMS/src
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build main.go

WORKDIR ${SOURCES}
CMD ${SOURCES}/main
# CMD [ "./main"]
# Entry point for executable(Giny which port to expose
EXPOSE 8080

#--------------------------
# FROM
# RUN commands
# ADD(optional)
# WORKDIR
# CMD