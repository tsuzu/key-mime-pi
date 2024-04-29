# Key Mime Pi

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](LICENSE)

## Difference from original

[The awesome original project](https://github.com/mtlynch/key-mime-pi) was deprecated and implemented in Python.
The goal of this project is to re-implement the original in Go and to keep it maintained.

If you are interested in the following features, check [TinyPilot](https://github.com/mtlynch/tinypilot).

* Video capture
* Support for OS- and browser-captured keystrokes (e.g., Ctrl+Alt+Del, Ctrl+W)
* Better stability

## Overview

Use your Raspberry Pi as a remote-controlled keyboard that accepts keystrokes through a web browser.

[![Key Mime Pi screenshot](https://raw.githubusercontent.com/mtlynch/key-mime-pi/master/screenshot.png)](https://raw.githubusercontent.com/mtlynch/key-mime-pi/master/screenshot.png)

## Compatibility

* Raspberry Pi Zero W
* Raspberry Pi 4 (Not Tested)

## Pre-requisites

* Raspberry Pi OS Stretch or later

## Quick Start

To begin, enable USB gadget support on the Pi by running the following commands:

```bash
sudo ./enable-usb-hid
sudo reboot
```

When the Pi reboots, run Key Mime Pi with the following commands:

Download binaries from [the latest release](https://github.com/tsuzu/key-mime-pi/releases/latest) and execute it as root.

example:
```console
$ wget https://github.com/tsuzu/key-mime-pi/releases/download/v0.0.1/key-mime-pi_Linux_armv6.tar.gz
...
$ tar xf key-mime-pi_Linux_armv6.tar.gz
$ sudo ./key-mime-pi 
2024/04/29 07:34:57 Listening on port :8000
```

Key Mime Pi will be running in your browser at:

* [http://raspberrypi:8000/](http://raspberrypi:8000/)


## Options(TODO: Fix)

Key Mime Pi accepts various options through environment variables:

| Environment Variable | Default      | Description |
|----------------------|--------------|-------------|
| `HOST`               | `0.0.0.0`    | Network interface to listen for incoming connections. |
| `PORT`               | `8000`       | HTTP port to listen for incoming connections. |
| `HID_PATH`           | `/dev/hidg0` | Path to keyboard HID interface. |
