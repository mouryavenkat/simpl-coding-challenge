# Get base centos image
FROM centos:7

#Install epel-release, debug tools, nginx, supervisor, boto packages.
#Create required directories
RUN \
  yum install -y epel-release bison python-setuptools bzip2 wget make gcc gcc-c++ zlib-devel git lsof && \
  easy_install supervisor &&  \
  mkdir -p /logs/manthan /etc/supervisord.d  && \
  yum clean all && \
  rm -f /etc/localtime && \
  ln -s /usr/share/zoneinfo/Asia/Kolkata /etc/localtime

#Install Go
RUN \
  cd /tmp && \
  wget https://storage.googleapis.com/golang/go1.12.6.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz && \
  ln -s /usr/local/go/bin/go /bin/go && \
  ln -s /usr/local/go/bin/gofmt /bin/gofmt

ENV GO111MODULE=auto CGO_ENABLED=1
ENV GODEBUG="madvdontneed=1"

EXPOSE 80

WORKDIR /usr/local/simpl/simpl-coding-challenge

COPY ./go.mod .

RUN go mod download

#Copy source directory
COPY ./ .
