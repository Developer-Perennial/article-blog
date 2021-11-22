FROM golang

WORKDIR /go/src/article-blog

COPY . .

RUN go get -d -v ./... &&\
    go install -v ./... &&\
    cp deployment/config.yaml /etc/config.yaml &&\
    rm -rf /go/src/article-blog

EXPOSE 8080

CMD ["article-blog", "-c", "/etc/config.yaml", "-m", "true"]