# build stage
FROM golang:alpine as build-stage
ARG enviroment
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ether_todo_service cmd/main.go || echo "🔥 go build failed"

# production stage
FROM golang:alpine as production-stage
COPY --from=build-stage /app/ether_todo_service ./ether_todo_service
EXPOSE 21000
CMD ["./ether_todo_service"]
