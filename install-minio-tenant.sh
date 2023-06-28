#!/bin/sh

helm upgrade --install \
  --namespace minio \
  --create-namespace \
  minio tenant --repo https://operator.min.io/ \
  --values ./minio-tenant.values.yml
