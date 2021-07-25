FROM golang:1.16
RUN mkdir -p /go/src/github.com/ho2ri2s/golang-dialy
# コンテナログイン時のディレクトリ指定
WORKDIR  /go/src/github.com/ho2ri2s/golang-dialy
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/ho2ri2s/golang-dialy