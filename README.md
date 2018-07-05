# Sidecar Support Application

```         
 _______ __     __         ______             
|     __|__|.--|  |.-----.|      |.---.-.----.
|__     |  ||  _  ||  -__||   ---||  _  |   _|
|_______|__||_____||_____||______||___._|__|  
```

# Overview

Sidecar is a stand-alone utility...

## Sidecar Responsibilities

Start/stop/restart containers.  Usually runs as a service from systemd on coreos.

### Primary Responsibilities

* read container configuration
* insure that containers meet config
* repeat

### Secondary Responsibilities

* clean logs
* bounce containers
* rolling updates

## Installation

### Systemd Service File

###### darryl.west | 2018.07.05

