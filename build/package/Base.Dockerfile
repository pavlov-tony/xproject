FROM golang:1.10-alpine

# Stuff needed for dep in alpine
RUN apk --update add alpine-sdk && \
  rm -rf /var/lib/apt/lists/* && \
  rm /var/cache/apk/*

# install dep
RUN curl -fsSL -o /usr/local/bin/dep \
  https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
  chmod +x /usr/local/bin/dep

# Dev dependencies
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get github.com/jstemmer/go-junit-report
RUN go get github.com/alecthomas/gometalinter && \
  gometalinter --install >/dev/null
  
WORKDIR /go/src/github.com/radisvaliullin/cloudbilling/cmd/cloudbilling