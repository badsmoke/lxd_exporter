# lxd_exporter



LXD metrics exporter for Prometheus. Improved version of [nieltg/lxdexporter](https://github.com/nieltg/lxd_exporter).

## Usage

Download latest docker image of this exporter from [the release page](https://github.com/nieltg/lxd_exporter/releases).

run as container
```
docker run -d -e LXD_DIR="/var/snap/lxd/common/lxd/" -v /var/snap/lxd/common/lxd/:/var/snap/lxd/common/lxd/ -p 9472:9472 lxd_exporter
```

docker-compose
```
version: "3"
services:
    lxd_exporter:
        image: badsmoke/lxd_exporter
        restart: always
        volumes:
            - '/var/snap/lxd/common/lxd/:/var/snap/lxd/common/lxd/'
        environment:
            - LXD_DIR="/var/snap/lxd/common/lxd/"
        ports:
            - 9472:9472

```

The exporter must have access to LXD socket which can be guided by:
- Specifying `LXD_SOCKET` environment variable to LXD socket path, or
- Specifying `LXD_DIR` environment variable to LXD socket's parent directory.

For more information, you can see the documentation from [Go LXD client library](https://godoc.org/github.com/lxc/lxd/client#ConnectLXDUnix).

## Hacking

docker build
```
docker build -t badsmoke/lxd_exporter:1.0.0 .

```
## License

[MIT](LICENSE).
