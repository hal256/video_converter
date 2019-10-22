FROM golang:1.13.1
RUN apt-get update && apt-get install xz-utils
WORKDIR /go/bin
RUN wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-i686-static.tar.xz && \
  tar xvf ffmpeg-release-i686-static.tar.xz
WORKDIR /go/bin/ffmpeg-4.2.1-i686-static
RUN cp ffmpeg /go/bin
RUN cp ffprobe /go/bin
COPY . .
RUN go get
RUN go build

CMD ["sh"]
