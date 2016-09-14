FROM golang:latest
ADD https://github.com/barockok/xyz_notif/archive/master.tar.gz ./artifact.tar.gz
RUN mkdir -p ./src/github.com/barockok/xyz_notif \
  && tar xf ./artifact.tar.gz -C ./src/github.com/barockok/xyz_notif --strip-components=1 \
  && rm ./artifact.tar.gz
WORKDIR ./src/github.com/barockok/xyz_notif
RUN go build github.com/barockok/xyz_notif
RUN cd cmd/xyz_thrift && go get ./...
RUN cd cmd/xyz_http && go get ./...
RUN go install github.com/barockok/xyz_notif/cmd/xyz_thrift
RUN go install github.com/barockok/xyz_notif/cmd/xyz_http
EXPOSE 3000
ENV PORT 3000
ENV PROTOCOL http
RUN echo "#!/bin/sh\nif [ \"\$PROTOCOL\" = \"http\" ]; then xyz_http; else xyz_thrift -addr localhost:\$PORT; fi" > ./entry.sh \
  && chmod +x ./entry.sh
CMD ["./entry.sh"]