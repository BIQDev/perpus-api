FROM golang:1.13-alpine
ARG BIQ_UID=1000
ARG BIQ_GID=1000
ENV SERVICE_NAME perpus-api
ENV GOPRIVATE gitlab.com/BIQDev/*

COPY build/assets/ghz /usr/local/bin
COPY build/assets/repositories /etc/apk/repositories

RUN chmod 644 /etc/apk/repositories && \
    apk update && apk add --no-cache sudo git
RUN apk add --no-cache make tzdata && \
    chmod a+x /usr/local/bin/ghz

COPY build/assets/sudoers /etc/sudoers

RUN chmod 0400 /etc/sudoers && \
    addgroup biqdev -g ${BIQ_GID} && \
    addgroup sudo && \
    adduser -h /home/biqdev -s /bin/sh -u ${BIQ_UID} -G users -D biqdev && \
    adduser biqdev biqdev && \
    adduser biqdev sudo

USER biqdev
WORKDIR "/go/src/github.com/BIQDev/${SERVICE_NAME}"

#COPY --chown=biqdev:users .netrc /tmp/.netrc

RUN sudo chown -R biqdev:biqdev /go && \
#    cp /tmp/.netrc ~/.netrc && \
#    chmod 600 ~/.netrc && \
    export GO111MODULE=off && \
    echo "Installing additional tools, please wait..." && \
    go get -u golang.org/x/lint/golint && echo "- Success installing golint" && \
    go get github.com/githubnemo/CompileDaemon && echo "- Success installing CompileDaemon" && \
    unset GO111MODULE


ENTRYPOINT ["sh", "./build/dev/entrypoint.sh"]
