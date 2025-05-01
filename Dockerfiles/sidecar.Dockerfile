ARG BASE_IMG_TAG=timberio/vector:nightly-debian
FROM ${BASE_IMG_TAG}

COPY vector.yaml /etc/vector/vector.yaml
