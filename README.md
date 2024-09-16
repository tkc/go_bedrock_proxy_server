# AWS Bedrock Proxy Server

This project provides a proxy server that sends signed requests to AWS Bedrock. It reads the AWS credentials from a YAML configuration file, signs the request using AWS Signature Version 4, and sends the signed request to the Bedrock API.

## Setup Configuration

Create a config.yaml file in the root of your project directory with the following structure:

```yaml
region: "us-west-2"
access_key_id: "YOUR_ACCESS_KEY"
secret_access_key: "YOUR_SECRET_ACCESS_KEY"
model_id: "YOUR_MODEL_ID"
```

## Run the Application

```bash
go run main.go
```

