ARG IMAGE
FROM ${IMAGE}
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends \
      curl \
      redis-tools && \
    rm -rf /var/lib/apt/lists/*
