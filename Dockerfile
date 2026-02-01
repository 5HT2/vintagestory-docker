FROM alpine:latest
ENV VERSION "1.19.4"
ENV ARCHIVE "vs_server_linux-x64_${VERSION}.tar.gz"

RUN mkdir /vintagestory \
 && apk --no-cache add ca-certificates wget \
 && wget "https://cdn.vintagestory.at/gamefiles/stable/${ARCHIVE}" \
 && tar xf "${ARCHIVE}" --directory vintagestory \
 && rm "${ARCHIVE}"
ADD . /vintagestory

FROM mono:latest
RUN mkdir /vintagestory \
 && mkdir /vintagestory-files
WORKDIR /vintagestory
COPY --from=0 /vintagestory .
RUN cert-sync /etc/ssl/certs/ca-certificates.crt
CMD [ "mono",  "./VintagestoryServer" ]
