# Orvibo S20 REST API

Provides a REST Api for S20 Orvibo wifi plugs 

# How to use

Just run in a network with S20 Orvibo wifi plugs. The server will automatically discover all of them.
Server is listening in port 8000. If you want to use the docker container you need to attach the container
to the host network:

``` bash
docker run --network host kdihalas/orvibo-s20-http-server
```

## Get a list of all discovered devices

The server will return a json with all discovered devices.


``` bash

curl http://localhost:8000/devices
...
...
{
  "ZZZNnnnnZZZ": {
    "ID": 1,
    "Name": "Socket ZZZNnnnnZZZ",
    "DeviceType": 0,
    "IP": {
      "IP": "10.10.10.10",
      "Port": 10000,
      "Zone": ""
    },
    "MACAddress": "ZZZNnnnnZZZ",
    "Subscribed": true,
    "Queried": true,
    "State": true,
    "RFSwitches": {},
    "LastIRMessage": "",
    "LastMessage": "6864002a716100ZZZNnnnnZZZ2020202020206a5c5123cfac202020202020534f433030355ad06ebc01"
  }
}

```

## Switch a plug to on/off

To switch a device on/off just query the switch endpoint with the device mac address and the state.
If the response is 200 then the device action has been submited successfully. 

``` bash

# Turn on
curl http://localhost:8000/switch?state=on&mac=ZZZNnnnnZZZ
*   Trying ::1:8000...
* Connected to localhost (::1) port 8000 (#0)
> GET /switch?state=on&mac=ZZZNnnnnZZZ HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.75.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Wed, 14 Apr 2021 11:55:28 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact


# Turn off
curl http://localhost:8000/switch?state=off&mac=ZZZNnnnnZZZ
*   Trying ::1:8000...
* Connected to localhost (::1) port 8000 (#0)
> GET /switch?state=off&mac=ZZZNnnnnZZZ HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.75.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Wed, 14 Apr 2021 11:55:28 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact

```