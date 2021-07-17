## go envt
FROM golang:latest

# setting working dirrectory
WORKDIR /go/src/go-blockchains
# copying programs in working dir.
COPY ./src .

# checking go version to ensure everything is good
RUN go version

CMD ["bash"]
