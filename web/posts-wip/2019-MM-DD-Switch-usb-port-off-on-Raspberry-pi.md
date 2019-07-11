# Switch USB ports on/off on a Raspberry Pi

I helped a friend co-worker to automate a water pump.

Find usb in `/sys/devices/platform/soc/`

ls /sys/devices/platform/soc/ | grep usb

sudo sh -c "echo 0 > /sys/devices/platform/soc/20980000.usb/buspower"
