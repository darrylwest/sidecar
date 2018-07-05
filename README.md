# Sidecar Support Application

```         
   _____ _     __     ______                   
  / ___/(_)___/ /__  / ____/___ ______
  \__ \/ / __  / _ \/ /   / __ `/ ___/
 ___/ / / /_/ /  __/ /___/ /_/ / /     
/____/_/\__,_/\___/\____/\__,_/_/       
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

