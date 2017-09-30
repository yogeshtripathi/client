# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.9
USER nobody
RUN mkdir -p /go/src/client
WORKDIR /go/src/client
COPY . /go/src/client
RUN go-wrapper download && go-wrapper install
CMD ["main"]
ENV PORT 9090
EXPOSE 9090
