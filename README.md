# Simple UDP generator

This is just for sending udp packet to make load.
Now under developing....

## Required

 * pcap library

## Usage (CLI mode)

```
#sudo go-udpgen send --help
go-udpgen send can be used to send UDP packates from CLI

Usage:
  go-udpgen send [Interface Name] [flags]
  
  Examples:
  $ go-udpgen send eth0 --dst-ip 10.10.10.10
  
  Flags:
        --concurrency int   The number of goroutines to use
        --count uint        The number of packets to be send (default 1)
        --dst-eth string    Dest mac address (default "00:00:00:00:00:02")
        --dst-ip string     Dest IP address (default "10.0.40.2")
        --dst-port string   UDP dest port (default "9999")
        --src-eth string    Source mac address (default "00:00:00:00:00:01")
        --src-ip string     Source IP address (default "10.0.40.1")
        --src-port string   UDP source port (default "9999")
        --time uint         seconds which keeps sending packtes
```

## Usage (WEB mode)

```
sudo go-udpgen server --port 9000
```

See help on detail.
