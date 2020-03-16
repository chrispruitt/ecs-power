NAME=ecs-power
VERSION=latest
DATE=`date +"%Y%m%d_%H%M%S"`
TEST_JSON='{"Power": "STATUS"}'
LAMBDA_NAME=ecs-power

clean:
	rm -rf dist

build: clean
	goreleaser --skip-publish --rm-dist --snapshot

zip: build
	cd dist/ecs-power_linux_amd64 && zip ../${NAME}.zip ${NAME}
	
updateLambda: zip
	aws lambda update-function-code --function-name ${LAMBDA_NAME} --zip-file fileb://${pwd}dist/${NAME}.zip

invoke:
	mkdir -p lambda_output
	aws lambda invoke \
		--function-name "${LAMBDA_NAME}" \
		--log-type "Tail" \
		--payload $(TEST_JSON) lambda_output/$(DATE).log \
		| jq -r '.LogResult' | base64 -D
