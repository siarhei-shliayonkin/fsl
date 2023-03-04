FROM alpine:3.17

EXPOSE 8081
COPY ./bin/fsl /usr/local/bin
ENTRYPOINT ["/usr/local/bin/fsl"]


FROM golang:1.18-alpine AS buildenv
WORKDIR /src
ADD . /src
RUN go mod download
RUN go build -o fsl cmd

# RUN chmod +x integritySum

FROM alpine:latest
COPY --from=buildenv /src/fsl /usr/local/bin

ENTRYPOINT ["/usr/local/bin/fsl"]
