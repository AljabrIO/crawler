# build stage
FROM golang:1.11-alpine AS build-env
RUN apk add -U git
ADD . /src
RUN cd /src && go get github.com/jackdanger/collectlinks && go build -o crawler

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/crawler /app/
RUN apk add -U ca-certificates
ENTRYPOINT [ "/app/crawler" ]
CMD [ ]
