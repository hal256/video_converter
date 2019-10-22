
docker build --rm -t video_convert:latest . && docker run -it -v $(pwd)/app:/go/src/app video_convert:latest
