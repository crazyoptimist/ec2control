# START STOP AWS EC2 INSTANCES

With AWS GO SDK

### Usage

Build it

```bash
make build
```

Run it

```bash
AWS_REGION=us-east-1 ./bin/ec2 -s START -i <your-ec2-id>
AWS_REGION=us-east-1 ./bin/ec2 -s STOP -i <your-ec2-id>
```

### License

MIT
