FROM golang
WORKDIR /dir-creation
RUN go mod init dir-creation
RUN mkdir cmd/
CMD ["/bin/bash"]
