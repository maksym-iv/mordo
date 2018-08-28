set -e

if [ -z $1 ]; then
  echo "Platform is not defined"
  echo "Supported platfoms: aws"
  echo "Usage: build_linux.sh aws"
  exit 1
fi

PLATFORM=$(echo $1 | tr "A-Z" "a-z")

docker build -t mordo -f scripts/Dockerfile --build-arg PLATFORM=${PLATFORM} ./
docker run --rm -t -v `pwd`/pkg:/pkg/build mordo sh -c "rm -rf /pkg/build/* && cp -R mordo lib /pkg/build/"

