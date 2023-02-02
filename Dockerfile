FROM alpine:3.17

EXPOSE 8081
COPY ./bin/fsl /usr/local/bin
ENTRYPOINT ["/usr/local/bin/fsl"]
