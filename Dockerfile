FROM alpine
COPY ./kube /
ENTRYPOINT /kube