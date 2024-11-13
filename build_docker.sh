REPOSITORY_URI=$1
VERSION=$2

docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm64/v8 \
  -t $REPOSITORY_URI:latest \
  -t $REPOSITORY_URI:$VERSION \
  --push \
  -f ./Dockerfile.cosmovisor \
  .