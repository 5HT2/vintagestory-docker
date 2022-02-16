FROM alpine:latest
RUN mkdir /vintagestory \
  && apk --no-cache add ca-certificates wget \
  && wget "https://cdn.vintagestory.at/gamefiles/stable/vs_archive_1.16.1.tar.gz" \
  && tar xf "vs_archive_1.16.1.tar.gz"
ADD . /vintagestory

FROM mono:latest
RUN mkdir /vintagestory \
  && mkdir /vintagestory-files
WORKDIR /vintagestory
COPY --from=0 /vintagestory .
RUN cert-sync /etc/ssl/certs/ca-certificates.crt
CMD [ "mono",  "./VintagestoryServer.exe" ]
