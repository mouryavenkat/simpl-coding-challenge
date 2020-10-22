# Get base centos image
FROM centos:7

RUN \
  yum install -y wget make gcc gcc-c++ git

#Install Go
RUN \
  cd /tmp && \
  wget https://storage.googleapis.com/golang/go1.12.6.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz && \
  ln -s /usr/local/go/bin/go /bin/go && \
  ln -s /usr/local/go/bin/gofmt /bin/gofmt

#Set environment variables
ENV PATH=$PATH:/usr/local/go/bin:/usr/local/simpl/simpl-coding-challenge/bin
ENV GO111MODULE=auto CGO_ENABLED=1
ENV GODEBUG="madvdontneed=1"

EXPOSE 80

WORKDIR /usr/local/simpl/simpl-coding-challenge

COPY ./go.mod .

RUN go mod download

#Copy source directory
COPY ./ .
