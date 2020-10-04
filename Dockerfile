FROM scratch

COPY gopath/bin/gosoma /gosoma

ENTRYPOINT ["/gosoma"]
