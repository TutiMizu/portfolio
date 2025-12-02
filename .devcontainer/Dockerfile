# Goの最新バージョンを使う
FROM golang:1.23-bookworm

# 開発に便利なツール（Gitなど）を入れておく
# Ubuntuベース(bookworm)なので apt が使える
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# コンテナ内の作業場所
WORKDIR /workspace
