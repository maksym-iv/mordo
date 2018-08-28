## Release build for client
set -e


RELEASE_DIR="release"


rm -rf ${RELEASE_DIR}

if ! [[ -d ${RELEASE_DIR} ]]; then
  mkdir ${RELEASE_DIR}
fi

# build Mordo binary
sh scripts/build_linux.sh aws

# Copy Mordo binary files
rsync -avpz --delete \
  pkg/ \
  ${RELEASE_DIR}/pkg

# Copy Terraform resources
rsync -avpz --delete \
  --exclude-from scripts/release/rsyncignore.txt \
  terraform-release/ \
  ${RELEASE_DIR}/terraform

# Copy docs
rsync -avpz --delete \
  --exclude-from scripts/release/rsyncignore.txt \
  doc/README.md \
  ${RELEASE_DIR}/
rsync -avpz --delete \
  --exclude-from scripts/release/rsyncignore.txt \
  doc/setup_AWS.md \
  ${RELEASE_DIR}/
rsync -avpz --delete \
  --exclude-from scripts/release/rsyncignore.txt \
  doc/setup_IAM.json \
  ${RELEASE_DIR}/

tar cfvpz mordo-release.tar.gz release
zip -r mordo-release.zip release