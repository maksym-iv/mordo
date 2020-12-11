FROM golang:latest

ARG PLATFORM
ARG VIPS_VERSION=8.10.2
ARG APP_ROOT=/app

RUN echo 'alias ll="ls -la"' >> /root/.bashrc
RUN apt-get update

RUN apt-get install -y automake build-essential git \
  libglib2.0-dev libjpeg62-turbo-dev libpng-dev \
  gtk-doc-tools libexpat1-dev libgif-dev zip

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /tmp/build
RUN wget -c https://github.com/libvips/libvips/releases/download/v${VIPS_VERSION}/vips-${VIPS_VERSION}.tar.gz
RUN tar xfvpz vips-${VIPS_VERSION}.tar.gz
WORKDIR /tmp/build/vips-${VIPS_VERSION}
RUN ./configure
RUN make
RUN ldconfig
RUN rm -rf /tmp/build/vips-${VIPS_VERSION}.tar.gz

# Build lambda deps
WORKDIR ${APP_ROOT}
ADD ./go.mod ${APP_ROOT}
ADD ./go.sum ${APP_ROOT}
RUN go mod download

ADD ./ /app

# Build more lambda deps
WORKDIR ${APP_ROOT}/function/${PLATFORM}

ENV CGO_ENABLED=1

RUN mkdir lib
RUN cp /usr/lib/x86_64-linux-gnu/libpng* ./lib/
RUN cp /usr/lib/x86_64-linux-gnu/libjpeg.* ./lib/
RUN cp /usr/lib/x86_64-linux-gnu/libgif.* ./lib/
RUN cp -R /tmp/build/vips-${VIPS_VERSION}/libvips/.libs/* ./lib/

ENV LD_LIBRARY_PATH="./lib"
ENV PKG_CONFIG_PATH="/tmp/build/vips-${VIPS_VERSION}"
ENV CGO_CFLAGS="-I/tmp/build/vips-${VIPS_VERSION}/libvips/include"
ENV CGO_LDFLAGS="-L/tmp/build/vips-${VIPS_VERSION}/libvips/.libs -lvips"

# Build lambda
RUN go build -a -tags release -o mordo
RUN ldd mordo
