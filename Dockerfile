FROM golang:1.14
WORKDIR /app
COPY . .
EXPOSE 8080
RUN go build
ENTRYPOINT ./helloworld
