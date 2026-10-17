FROM golang:1.24 AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /media-saver .

FROM python:3.12-slim
RUN apt-get update && apt-get install -y --no-install-recommends ffmpeg rclone ca-certificates \
 && rm -rf /var/lib/apt/lists/* \
 && pip install --no-cache-dir -U yt-dlp gallery-dl
COPY --from=build /media-saver /usr/local/bin/media-saver
ENV RCLONE_CONFIG=/config/rclone.conf
EXPOSE 8080
CMD ["media-saver"]
