set -e

sh scripts/build_linux.sh aws 1>/dev/null


cat doc/apigw_event/blank.json |  docker run --rm -i -v "$PWD"/pkg:/var/task --env-file=.aws_creds --env-file=.envlocal --env-file=.env lambci/lambda:go1.x mordo 2>&1