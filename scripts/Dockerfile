FROM golang:latest

ARG PLATFORM

RUN echo 'alias ll="ls -la"' >> /root/.bashrc
RUN apt-get update

RUN apt-get install -y automake build-essential git \
  libglib2.0-dev libjpeg62-turbo-dev libpng-dev \
  gtk-doc-tools libexpat1-dev libgif-dev zip

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /tmp/build
RUN wget -c https://github.com/jcupitt/libvips/releases/download/v8.6.5/vips-8.6.5.tar.gz
RUN tar xfvpz vips-8.6.5.tar.gz
WORKDIR /tmp/build/vips-8.6.5
RUN ./configure
RUN make
# RUN make install
RUN ldconfig
RUN rm -rf /tmp/build/vips-8.6.5.tar.gz

# Build lambda
ADD ./Gopkg.toml /go/src/github.com/xmackex/mordo/Gopkg.toml
ADD ./Gopkg.lock /go/src/github.com/xmackex/mordo/Gopkg.lock
WORKDIR /go/src/github.com/xmackex/mordo
RUN dep ensure --vendor-only

ADD ./ /go/src/github.com/xmackex/mordo

# For function
WORKDIR /go/src/github.com/xmackex/mordo/function/${PLATFORM}

ENV CGO_ENABLED=1

RUN mkdir lib
RUN cp /usr/lib/x86_64-linux-gnu/libpng* lib/
RUN cp /usr/lib/x86_64-linux-gnu/libjpeg.* lib/
RUN cp /usr/lib/x86_64-linux-gnu/libgif.* lib/
RUN cp -R /tmp/build/vips-8.6.5/libvips/.libs/* ./lib/

ENV LD_LIBRARY_PATH="./lib"
ENV PKG_CONFIG_PATH="/tmp/build/vips-8.6.5"
ENV CGO_CFLAGS="-I/tmp/build/vips-8.6.5/libvips/include"
ENV CGO_LDFLAGS="-L/tmp/build/vips-8.6.5/libvips/.libs -lvips"

RUN go build -a -tags release -o mordo
RUN ldd mordo

# CC=/usr/local/musl/bin/musl-gcc go build -a -ldflags '-linkmode external -extldflags "-static"'

# RUN zip -r lambda_deploy lambda lib
