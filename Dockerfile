FROM alpine
COPY ./app /usr/local/bin
ENTRYPOINT /usr/local/bin