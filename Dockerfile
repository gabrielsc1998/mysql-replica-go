FROM golang:1.22

WORKDIR /go/src/
COPY . .

CMD ["tail", "-f", "/dev/null"]

# ENTRYPOINT [ "sh", "./.docker/entrypoint.sh" ]