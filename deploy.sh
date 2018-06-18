#!/bin/bash
FUNCTION_NAME="dexpenses-classify"
BUCKET="dexpenses-classify-code"
FILENAME="deployment-package.zip"
OUTPUT="out"
PACKAGE="main"
rm -r ${OUTPUT} || true
mkdir ${OUTPUT}
echo Creating package...
go build -o ${OUTPUT}/${PACKAGE} || exit 1
pushd ${OUTPUT}
if zip ${FILENAME} ${PACKAGE}; then
    popd
else
    popd
    exit 1
fi
zip ${FILENAME} ${PACKAGE} || popd && exit 1
popd
echo Uploading to S3...
aws s3 cp --acl public-read ./${OUTPUT}/${FILENAME} s3://${BUCKET}/${FILENAME} || exit 1
echo Updating function...
aws lambda update-function-code --function-name ${FUNCTION_NAME} --s3-bucket ${BUCKET} --s3-key ${FILENAME}