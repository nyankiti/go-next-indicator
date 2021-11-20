FROM golang:1.17

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apt-get update \
  && go mod download \
  && apt-get install -y mariadb-client


COPY . .

#CMD ["go", "run", "main.go"]
CMD ["sh", "./start_app.sh"]