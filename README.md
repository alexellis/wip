## WIP - A sample Go Function for OpenFaaS

This is a work in progress (WIP) OpenFaaS function written with the Go template.

## Build the code

```bash
sudo -E arkade get faas-cli --path /usr/local/bin/
faas-cli template pull golang-middleware

curl -sLS https://get.docker.com | sudo sh
sudo useradd -aG docker $USER
newgrp docker
sudo systemctl start docker

faas-cli build
```
