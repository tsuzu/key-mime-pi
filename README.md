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

TODO: Fix

```
python3 -m venv venv
. venv/bin/activate
pip install --requirement requirements.txt
PORT=8000 ./app/main.py
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
