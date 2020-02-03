NAME=ecs-power
VERSION=latest
DATE=`date +"%Y%m%d_%H%M%S"`
TEST_JSON='{ "Cluster": "tst", "Power": "OFF"}'
LAMBDA_NAME=ecs-power
DOCKER_ARGS=--name $(NAME) \
	--rm \
	-v "`pwd`/build":/var/task \
	-e DEBUG=true \
	-e AWS_ACCESS_KEY_ID \
	-e AWS_SECRET_ACCESS_KEY \
	-e AWS_SESSION_TOKEN \
	lambci/lambda:go1.x $(NAME)

clean:
	rm -rf dist

build: clean
	goreleaser --skip-publish --rm-dist --snapshot

zip: build
	cd dist/ecs-power_linux_amd64 && zip ../${NAME}.zip ${NAME}
	
updateLambda:
	aws lambda update-function-code --function-name ${LAMBDA_NAME} --zip-file fileb://${pwd}dist/${NAME}.zip

invoke:
	mkdir -p lambda_output
	aws lambda invoke \
		--function-name "${LAMBDA_NAME}" \
		--log-type "Tail" \
		--payload $(TEST_JSON) lambda_output/$(DATE).log \
		| jq -r '.LogResult' | base64 -D
