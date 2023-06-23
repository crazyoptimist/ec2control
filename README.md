# EC2CONTROL

This is a CLI application that can start/stop AWS EC2 instances on demand.

### Build

```bash
make build
```

It will generate binaries for linux and windows inside the `./bin` directory.

### Configuration

You need to create a `config.toml` file. The app will print the template if it does not exist.

Config file path can be the same directory as the binary, or:

- Linux

`$HOME/.config/ec2/config.toml`

- Windows

`%AppData%\ec2\toml`

### Usage

- Linux

```bash
./ec2control -start
```

- Windows

```bash
ec2control -start
```

### License

**MIT**

Copyright 2023 crazyoptimist

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
