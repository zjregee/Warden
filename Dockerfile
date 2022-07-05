FROM golang
MAINTAINER  yrj
WORKDIR /go/src/
COPY . .
EXPOSE 8000
CMD ["/bin/bash", "/go/src/script/build.sh"]