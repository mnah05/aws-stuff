import * as cdk from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import * as apigw from 'aws-cdk-lib/aws-apigateway'

import * as lambda from 'aws-cdk-lib/aws-lambda'
export class AwsLambdaGoStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'AwsLambdaGoQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });

    const fn = new lambda.Function(this, 'GoHandler', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      handler: 'bootstrap',
      code: lambda.Code.fromAsset('lambdas'),
      architecture: lambda.Architecture.ARM_64
    })
    const api = new apigw.LambdaRestApi(this, 'GoApi', {
      handler: fn,
      proxy: false,
      defaultCorsPreflightOptions: {
        allowOrigins: apigw.Cors.ALL_ORIGINS,
        allowMethods: apigw.Cors.ALL_METHODS,
      },
    });
    api.root.addResource('hi').addMethod('GET', new apigw.LambdaIntegration(fn));
  }
}
