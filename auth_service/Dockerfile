FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Скопируем остальные файлы проекта
COPY . .

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM scratch AS final

WORKDIR /

COPY --from=build /bin/main /main

EXPOSE 8081

CMD [ "/main" ]