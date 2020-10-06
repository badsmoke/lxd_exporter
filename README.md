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
            - "/var/snap/lxd/common/lxd/:/var/snap/lxd/common/lxd/"
        environment:
            - "LXD_DIR=/var/snap/lxd/common/lxd/"
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

Output hostname:9472/metrics

```
# HELP lxd_instance_cpu_usage instanceCPU Usage in Seconds
# TYPE lxd_instance_cpu_usage gauge
lxd_instance_cpu_usage{instance_name="c1",instance_type="container"} 1.814131565e+10
lxd_instance_cpu_usage{instance_name="v1",instance_type="virtual-machine"} 6.695872399e+09
# HELP lxd_instance_mem_usage instanceMemory Usage
# TYPE lxd_instance_mem_usage gauge
lxd_instance_mem_usage{instance_name="c1",instance_type="container"} 3.48270592e+08
lxd_instance_mem_usage{instance_name="v1",instance_type="virtual-machine"} 2.99933696e+08
# HELP lxd_instance_mem_usage_peak instanceMemory Usage Peak
# TYPE lxd_instance_mem_usage_peak gauge
lxd_instance_mem_usage_peak{instance_name="c1",instance_type="container"} 5.81820416e+08
lxd_instance_mem_usage_peak{instance_name="v1",instance_type="virtual-machine"} 2.9556736e+08
# HELP lxd_instance_network_usage instanceNetwork Usage
# TYPE lxd_instance_network_usage gauge
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="eth0",operation="BytesReceived"} 767411
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="eth0",operation="BytesSent"} 22734
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="eth0",operation="PacketsReceived"} 418
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="eth0",operation="PacketsSent"} 231
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="lo",operation="BytesReceived"} 1848
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="lo",operation="BytesSent"} 1848
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="lo",operation="PacketsReceived"} 20
lxd_instance_network_usage{instance_name="c1",instance_type="container",interface="lo",operation="PacketsSent"} 20
lxd_instance_network_usage{instance_name="v1",instance_type="virtual-machine",interface="enp5s0",operation="BytesReceived"} 4237
lxd_instance_network_usage{instance_name="v1",instance_type="virtual-machine",interface="enp5s0",operation="BytesSent"} 4635
lxd_instance_network_usage{instance_name="v1",instance_type="virtual-machine",interface="enp5s0",operation="PacketsReceived"} 48
lxd_instance_network_usage{instance_name="v1",instance_type="virtual-machine",interface="enp5s0",operation="PacketsSent"} 53
# HELP lxd_instance_pid instancePID
# TYPE lxd_instance_pid gauge
lxd_instance_pid{instance_name="c1",instance_type="container"} 2.63507e+06
lxd_instance_pid{instance_name="v1",instance_type="virtual-machine"} 2.639881e+06
# HELP lxd_instance_process_count instancenumber of process Running
# TYPE lxd_instance_process_count gauge
lxd_instance_process_count{instance_name="c1",instance_type="container"} 51
lxd_instance_process_count{instance_name="v1",instance_type="virtual-machine"} 20
# HELP lxd_instance_running_status instanceRunning Status
# TYPE lxd_instance_running_status gauge
lxd_instance_running_status{instance_name="c1",instance_type="container"} 1
lxd_instance_running_status{instance_name="v1",instance_type="virtual-machine"} 1
# HELP lxd_instance_swap_usage instanceSwap Usage
# TYPE lxd_instance_swap_usage gauge
lxd_instance_swap_usage{instance_name="c1",instance_type="container"} 0
lxd_instance_swap_usage{instance_name="v1",instance_type="virtual-machine"} 0
# HELP lxd_instance_swap_usage_peak instanceSwap Usage Peak
# TYPE lxd_instance_swap_usage_peak gauge
lxd_instance_swap_usage_peak{instance_name="c1",instance_type="container"} 0
lxd_instance_swap_usage_peak{instance_name="v1",instance_type="virtual-machine"} 0

```



## License

[MIT](LICENSE).
