# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go-fuzz/ /go/src/github.com/dvyukov/go-fuzz/
RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz && go build
RUN cd /go/src/github.com/dvyukov/go-fuzz/go-fuzz && cp go-fuzz /usr/local/bin

# Copy source codes
COPY . /app
RUN cd /app && go build
RUN mv input-generator sonar.exe
RUN apt-get update && apt-get install -y zip
RUN zip fuzz.zip sonar.exe metadata

RUN cd /app && mv fuzz.zip /usr/local/bin
# RUN cd /app && mkdir corpus && cd corpus && echo "1" > init_corpus

# Expose port 8080 to the outside world
EXPOSE 19960 

# Command to run the executable
CMD ["go-fuzz", "--procs=1", "--bin=/usr/local/bin/fuzz.zip", "--timeout=300", "-v=5"]

