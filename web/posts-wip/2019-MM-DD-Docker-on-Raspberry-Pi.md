# Install Docker on Raspberry Pi

This is a script that I wrote a couple of months ago for my Raspberry Pi.

It installs:

 - Docker
 - Docker Compose (using [pip](https://pypi.org/project/pip/))

```bash
# Run with super user priviledges
#!/bin/bash

# Update packages
echo "Update packages"
apt-get update

# Remove old docker
echo "Remove old docker"
apt-get -y remove docker docker-engine docker.io containerd runc
rm -rf /var/lib/docker

# Install dependencies
echo "Install dependencies"
apt-get -y install \
     apt-transport-https \
     ca-certificates \
     curl \
     gnupg2 \
     software-properties-common

# Download docker scripts
echo "Download docker scripts"
curl -fsSL https://get.docker.com -o get-docker.sh

# Give executable right
echo "Give executable right"
chmod +x get-docker.sh

# Run Docker installer
echo "Run Docker installer"
sh get-docker.sh

# Add the pi user to be able to run docker
echo "Add the pi user to be able to run docker"
usermod -aG docker pi

# Start docker daemon at boot
echo "Start docker daemon at boot"
systemctl enable docker

# Clean up
echo "Clean up"
rm -rf get-docker.sh

# Docker compose
apt-get install -y python-pip
pip install docker-compose

# Reboot
echo "Please, reboot. (sudo reboot)"
```

Hope that helps!
