FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Build go from source
WORKDIR /
RUN git clone https://go.googlesource.com/go goroot
WORKDIR /goroot
RUN git checkout go1.23.5
WORKDIR /goroot/src
RUN ./all.bash
WORKDIR /
RUN rm -rf /usr/local/go
RUN mv /goroot/ /usr/local/go

RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg
RUN apt-get update && apt-get upgrade -y && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client-17

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/xo/xo@latest
USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
