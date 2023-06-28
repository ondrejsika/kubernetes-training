#!/bin/sh

helm upgrade --install \
  --namespace minio-operator \
  --create-namespace \
  minio-operator operator --repo https://operator.min.io/
