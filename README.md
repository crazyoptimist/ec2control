# EC2CONTROL

This CLI application start/stop AWS EC2 instances as needed.

## Get AWS access credentials

To use this application, you need to create AWS credentials:

https://docs.aws.amazon.com/keyspaces/latest/devguide/access.credentials.html

Your key pair needs to have below permissions at least:

```
"ec2:StartInstances"
"ec2:StopInstances"
```

## Usage

```bash
ec2control help
```
