#!/usr/bin/python

import subprocess

def shell_exec(cmd, shell=True):

    try:
        ret = subprocess.check_output(cmd,shell=shell)
    except subprocess.CalledProcessError as e:
        return e

    return ret


timeout = 10
while timeout:
    discovery = shell_exec("curl -s https://discovery.etcd.io/new?size=1")
    if "etcd.io" in discovery:
        break
    timeout -= 1

if "etcd.io" in discovery:
    discovery_http = discovery.replace("https", "http")

start_cmd = "docker run -d --name wharf-meta-etcd \
                -v /var/lib/etcd-wharf-meta:/var/lib/etcd-wharf-meta \
              --net=host --privileged --restart=on-failure \
                gcr.io/google-containers/etcd:3.1.11 etcd \
              --name=wharf-meta \
              --initial-advertise-peer-urls=http://127.0.0.1:2380 \
              --listen-peer-urls=http://127.0.0.1:2380 \
              --listen-client-urls=http://127.0.0.1:2379 \
              --advertise-client-urls=http://127.0.0.1:2379 \
              --discovery={discovery} \
              --data-dir=/var/lib/etcd-wharf-meta".format(discovery=discovery_http)

shell_exec(start_cmd)
