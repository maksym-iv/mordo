set -e

if [ -z $1 ]; then
  echo "Build env is not defined"
  echo "Supported envs: aws, go"
  echo "Usage: build_linux.sh aws/go aws"
  exit 1
fi

if [ -z $2 ]; then
  echo "Platform is not defined"
  echo "Supported platfoms: aws"
  echo "Usage: build_linux.sh aws/go aws"
  exit 1
fi

BUILD_ENV=$(echo $1 | tr "A-Z" "a-z")
PLATFORM=$(echo $2 | tr "A-Z" "a-z")

docker build -t mordo-${BUILD_ENV} -f scripts/Dockerfile-${BUILD_ENV} --build-arg PLATFORM=${PLATFORM} ./
docker run --rm -t -v `pwd`/pkg:/pkg/build mordo-${BUILD_ENV} sh -c "rm -rf /pkg/build/* && cp -R mordo lib /pkg/build/"
