# sample-circleci

[![CircleCI](https://circleci.com/gh/d-tsuji/sample-circleci.svg?style=svg)](https://app.circleci.com/pipelines/github/d-tsuji/sample-circleci)

This repository is intended to show how to configure CircleCI to run Go tests with DynamoDB Local.

I use the AWS CLI to manipulate DynamoDB. Therefore, I have installed the AWS CLI in [config.yml](https://github.com/d-tsuji/sample-circleci/blob/master/.circleci/config.yml).

If you do not use the CLI to set up DynamoDB, you do not need it.
