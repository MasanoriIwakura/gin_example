FROM golang:1.16.4

# gitのインストールはコマンドサンプル(元から入っている)
# 他のツール入れたらgit消すこと
RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  git && \
  apt-get -y clean && \
  rm -rf /var/lib/apt/lists/*

# 開発用ツール
RUN go get -u github.com/cosmtrek/air && \
    go get -u github.com/go-delve/delve/cmd/dlv

WORKDIR /app

COPY ./app ./

RUN go build
