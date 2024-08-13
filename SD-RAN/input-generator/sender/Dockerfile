# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app
COPY go-fuzz/ /go/src/github.com/dvyukov/go-fuzz/

# Copy source codes
COPY . /app
RUN cd /app && go build

RUN cd /app && mv input-generator /usr/local/bin
# RUN cd /app && mkdir corpus && cd corpus && echo "1" > init_corpus

# Expose port 8080 to the outside world
EXPOSE 19960 

# Command to run the executable
CMD ["input-generator"]

