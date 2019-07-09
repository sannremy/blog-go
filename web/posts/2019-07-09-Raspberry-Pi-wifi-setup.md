# Raspberry Pi + WiFi setup

_Tested on macOS Mojave._

Hardware:

 - Raspberry Pi 2 model B+.
 - USB WiFi adapter: TP-Link WN725N v2 (Based on the Realtek RTL8188EUS chipset - 8188eu drivers).

## Boot disk

### Format SD card (FAT32)

N = disk ID, e.g. `disk2`. It can be found in Disk Utility.

```bash
diskutil eraseDisk FAT32 RASPBERRYPI /dev/diskN
```
### Create Raspbian boot from image on SD card

Unmount the drive.

```bash
diskutil unmountDisk /dev/diskN
```

[Download](https://www.raspberrypi.org/downloads/raspbian/) an official image of Raspbian.

```bash
sudo dd bs=1m if=/path/to/raspbian.img of=/dev/rdiskN conv=sync
```
## Install Raspbian

Insert the SD card and turn the Raspberry Pi on.
Login: pi/raspberry

## Enable WiFi

```bash
sudo raspi-config
```
Select _Network_ to set up WiFi.

Enjoy! 🙂
