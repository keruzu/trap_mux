
BUILDARCH=x86_64
TARGET=trapmux
image = alpine:3.15
docker_tag_trapmux = kkearne/trapmux
container_trapmux = trapmux
configuration_path_trapmux = /Users/kellskearney/go/src/trapmux/tools


codebuild:
# Need to run the following first
# aws configure
	#aws cloudformation deploy --template-file tools/aws/codebuild_cfn.yml --stack-name trapmuxrpm --capabilities CAPABILITY_IAM
	aws cloudformation deploy --template-file tools/aws/codebuild_docker.yml --stack-name trapmuxdocker --capabilities CAPABILITY_IAM
	#aws cloudformation deploy --template-file tools/aws/codebuild_batch_cfn.yml --stack-name trapmuxbatchrpm --capabilities CAPABILITY_IAM --parameter-overrides StreamId=rpm BuildSpec=tools/aws/buildspec_batch_rpm.yml
	#aws cloudformation deploy --template-file tools/aws/codebuild_batch_cfn.yml --stack-name trapmuxbatchnopkg --capabilities CAPABILITY_IAM --parameter-overrides StreamId=nopkg BuildSpec=tools/aws/buildspec_batch_nopkg.yml CodeBuildImage=aws/codebuild/standard:5.0 

