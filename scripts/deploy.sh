#!/bin/bash
DEFAULT_OUTPUT="bin"
DEFAULT_PACKAGE="main"

WD=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P ) # pwd of this script
PROJECT_DIR="$(dirname "$WD")"
PROJECT_FOLDER_NAME="${PROJECT_DIR##*/}"
source ${PROJECT_DIR}/deployments/.aws-config || exit 1
if [ -z ${BUCKET} ]; then
    echo "BUCKET variable must be set in ../deployments/.aws-config"
    exit 1
fi
OUTPUT="$PROJECT_DIR/${OUTPUT:-$DEFAULT_OUTPUT}"

if [ -z ${FILENAME} ]; then
    FILENAME="$PROJECT_FOLDER_NAME-deployment-pkg.zip"
fi
if [ -z ${FUNCTION_NAME} ]; then
    FUNCTION_NAME=${PROJECT_FOLDER_NAME}
fi
echo "Using function name $FUNCTION_NAME"
PACKAGE="${PACKAGE:-$DEFAULT_PACKAGE}"
rm -r ${OUTPUT} || true
mkdir ${OUTPUT}

echo Creating package...
cd ${PROJECT_DIR}
go build -o ${OUTPUT}/${PACKAGE} cmd/${FUNCTION_NAME}/main.go || exit 1

zip -j ${OUTPUT}/${FILENAME} ${OUTPUT}/${PACKAGE} || exit 1
echo Uploading to S3...
aws s3 cp --acl public-read ${OUTPUT}/${FILENAME} s3://${BUCKET}/${FILENAME} || exit 1
echo Updating function...
aws lambda update-function-code --function-name ${FUNCTION_NAME} --s3-bucket ${BUCKET} --s3-key ${FILENAME}