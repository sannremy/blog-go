# Switch USB ports on/off on a Raspberry Pi

I helped a friend co-worker to automatically turn on/off a pump for watering plants.

Find usb in `/sys/devices/platform/soc/`

```bash
ls /sys/devices/platform/soc/ | grep usb
```

Off
```bash
sudo sh -c "echo 0 > /sys/devices/platform/soc/20980000.usb/buspower"
```

On
```bash
sudo sh -c "echo 1 > /sys/devices/platform/soc/20980000.usb/buspower"
```

Cron 9pm to 9:15pm
```bash
0 21 * * * sudo sh -c "echo 1 > /sys/devices/platform/soc/20980000.usb/buspower" >/dev/null 2>&1

15 21 * * * sudo sh -c "echo 0 > /sys/devices/platform/soc/20980000.usb/buspower" >/dev/null 2>&1
```
