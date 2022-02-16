FROM alpine:latest
ENV VERSION "1.16.3"

RUN mkdir /vintagestory \
  && apk --no-cache add ca-certificates wget \
  && wget "https://cdn.vintagestory.at/gamefiles/stable/vs_archive_${VERSION}.tar.gz" \
  && tar xf "vs_archive_${VERSION}.tar.gz"
ADD . /vintagestory

FROM mono:latest
RUN mkdir /vintagestory \
  && mkdir /vintagestory-files
WORKDIR /vintagestory
COPY --from=0 /vintagestory .
RUN cert-sync /etc/ssl/certs/ca-certificates.crt
CMD [ "mono",  "./VintagestoryServer.exe" ]
