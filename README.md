# AWS Lambda performance testing

An attempt to create a performance testing suite for AWS lamnbdas

## Overview

This project is a set of [AWS Lambdas](https://aws.amazon.com/lambda/) written in different languages to enable benchmarking each one.
The idea is they'll perform typical real world user cases rather than being a simple "Hello, World!" Lambda.


## Working locally

To run the stack locally you'll need:

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [GNU make](https://www.gnu.org/software/make/manual/make.html)

Once all the tools are correctly installed you can run the following command:

```sh
$ make sam-start
```

This will spin up local [API Gateway](https://aws.amazon.com/api-gateway/) on port `3000`.
To create the DynamoDB table run the following command:

```sh
$ make sam-create-table
```

There is also a Make target that can load some data in as well.
This is done by running:

```sh
$ make sam-data-load
```

If the commands were all successful you sould be able to get a result back from making a HTTP request to `localhost:3000/nodejs-810/football/english-premier-league/season-1718/result?HomeTeam=Arsenal&AwayTeam=Tottenham"`.
With [`curl`](https://curl.haxx.se/):

```sh
$ curl "localhost:3000/nodejs-810/football/english-premier-league/season-1718/result?HomeTeam=Arsenal&AwayTeam=Tottenham
```

With [HTTPie](https://httpie.org/):

```sh
$ http localhost:3000/nodejs-810/football/english-premier-league/season-1718/result HomeTeam=="Arsenal" AwayTeam=="Tottenham"
```

## Deploying the code

### Pre-reqs
First fork this repo. Once forked you'll need: 

* An [AWS account](https://aws.amazon.com/account/)
* A GitHub Personal Access Token, for more information see [Creating a personal access token for the command line](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) on the GitHub website.

### The AWS CodePipeline
In the folder [deployments/cloudformation](deployments/cloudformation/) folder there are four files:

1. **[master.yaml](deployments/cloudformation/master.yaml):** This is the master CloudFormation template file that will deploy the other stacks.
2. **[pipeline-roles.yaml](deployments/cloudformation/pipeline-roles.yaml):** The CloudFormation template used to create the roles needed to run the build CodeBuild, CodePipeline and deploy the Lambdas.
3. **[artifact-bucket.yaml](deployments/cloudformation/artifact-bucket.yaml):** The CloudFormation template used to create the Amazon S3 bucket for the build artifacts.
4. **[pipeline.yaml](deployments/cloudformation/pipeline.yaml):** The CloudFormation template used to create the AWS CodePipeline and the AWS CodeBuild steps to build and deploy the Lambdas.

To create the benchmark pipeline stack.

[<img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png">](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/new?stackName=lambda-benchmark&templateURL=https://s3.amazonaws.com/aws-lambda-benchmark/master.yaml)

Once the stack is complete, we will need to trigger the build to deploy the Lambdas.
The output of the stack will have a link to the AWS CodePipeline from where you can kick the build off.
