<!--
SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Condensed instructions for setting up podmand and kind on MacOS

These instructions work for Intel-based Macs only for now.

## Install podman and start the VM
Start with installing `podman`:
```
brew install podman
```
Once done, init and start the machine:
```
podman machine init --volume $HOME:$HOME
podman machine start
```

## Install system helper service
To avoid having to export `DOCKER_HOST` environment variable, you can install the MacOS
system helper service via the following commands:
```
sudo /usr/local/Cellar/podman/4.0.3/bin/podman-mac-helper install
podman machine stop; podman machine start
```
Note that this is a one-time setup. Once run, you won't need to do this again even if you re-create and re-modify
the `podman` machine via the following steps.

## Patch podman.service on the VM
At this point some changes have to be made to the podman machine to allow it
to operate with kind in rootless mode:

```
podman machine ssh
```

Once on the machine follow these steps:

```
sudo cp /usr/lib/systemd/user/podman.service /etc/systemd/user/
sudo vi /etc/systemd/user/podman.service
```

Add `Delegate=true` line under `[Service]` section:
```
...
[Service]
Delegate=true
...
```

## Patch IPv6 Tables on the VM
While still logged in to the podman machine run the following commands:

```
curl -O https://kojipkgs.fedoraproject.org/packages/podman/4.0.2/1.fc35/x86_64/podman-4.0.2-1.fc35.x86_64.rpm
sudo -i
```

Patch the IPv6 tables...

```
rpm-ostree override replace /home/core/podman-4.0.2-1.fc35.x86_64.rpm
echo ip6_tables > /etc/modules-load.d/ip6_tables.conf
systemctl reboot
```

## Install docker, kind, kubectl and helm
```
brew install docker kind kubectl helm
```

## Setup µONOS Cluster
```
export KIND_EXPERIMENTAL_PROVIDER=podman
cd build-tools/dev
./setup-cluster
```

Note that after following all these steps, you may need to start a separate terminal shell and/or you
may need to stop and restart the `podman` machine.

Best of luck... ;)