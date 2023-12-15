ARG GO_VERSION=1.21.5
ARG ALPINE_VERSION=3.17

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# set default ENV values (like timezone/etc etc.)
ENV TZ=UTC \
	DEBIAN_FRONTEND=noninteractive \
	LANG=en_US.UTF-8 \
	LC_ALL=C.UTF-8

# Setup user
RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Install packages and Go language server
RUN apk add -q --update --progress --no-cache git sudo openssh-client zsh
RUN go get -u -v golang.org/x/tools/cmd/gopls 2>&1

# install golang task runner
RUN go get -u -v github.com/go-task/task/cmd/task 2>&1

# install golang mage tool
RUN go get -u -v github.com/magefile/mage 2>&1

USER root
