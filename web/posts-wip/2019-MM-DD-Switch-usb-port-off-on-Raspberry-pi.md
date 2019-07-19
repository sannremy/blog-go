# Switch USB ports on/off on a Raspberry Pi

I helped a friend co-worker to automatically turn on/off a pump for watering plants.

Find usb in `/sys/devices/platform/soc/`

```bash
ls /sys/devices/platform/soc/ | grep usb
```

Off
```bash
sh -c "echo 0 > /sys/devices/platform/soc/20980000.usb/buspower"
```

On
```bash
sh -c "echo 1 > /sys/devices/platform/soc/20980000.usb/buspower"
```

Modify the root users crontab
```bash
sudo crontab -e
```

Cron 9pm to 9:15pm
```bash
0 21 * * * sh -c "echo 1 > /sys/devices/platform/soc/20980000.usb/buspower" >/dev/null 2>&1

15 21 * * * sh -c "echo 0 > /sys/devices/platform/soc/20980000.usb/buspower" >/dev/null 2>&1
```
