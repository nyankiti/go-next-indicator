FROM golang:1.17

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apt-get update \
  && go mod download \
  && apt-get install -y mariadb-client


COPY . .

# live loadingを可能にするため、airの導入
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

#CMD ["go", "run", "main.go"]
#CMD ["sh", "./start_app.sh"]
CMD ["air"]
