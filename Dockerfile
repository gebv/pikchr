FROM golang:1.14 as build

    WORKDIR /go/src/github.com/gebv/pikchr
    COPY go.mod .
    COPY go.sum .
    RUN go mod download
    COPY . .

    RUN make build-render-server

FROM heroku/heroku:18
    WORKDIR /app
    ENV HOME /app

    COPY --from=build /go/src/github.com/gebv/pikchr/bin/render-server .

    RUN useradd -m heroku
    USER heroku
    CMD /app/render-server
