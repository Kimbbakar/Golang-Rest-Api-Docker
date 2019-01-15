# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/kimbbakar/rest-api/
# Add the source code:
ADD . $SRC_DIR
RUN go get github.com/gorilla/mux 

# Build it:
RUN cd $SRC_DIR; go build -o rest-api; cp rest-api /app/
ENTRYPOINT ["./rest-api"]
