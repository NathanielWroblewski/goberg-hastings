FROM golang:1-alpine

COPY . /src
WORKDIR /src

CMD ["go", "run", "main.go", "-width=100", "-height=100"]
