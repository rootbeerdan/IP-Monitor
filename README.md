# IP-Monitor
Utility to monitor dual stack network status

## Purpose

This small utility is intended to be deployed to end-user devices on networks that rely on IPv6-only services, so they can be aware of current network conditions.

## Use

There are 3 colors that signify the status of the network:

- Green: IPv4 and IPv6 are reachable
- Yellow: IPv4 only, no publicly routed IPv6 services
- Red: No IPv4 or IPv6

Users will recieve a notification for any status changes.

## Configuration (conf.json)

coming soon

## Notes

This utility by default relies on access to Akamai IP services, including the following endpoints:

- http://ipv6.whatismyip.akamai.com
- http://whatismyip.akamai.com

Polling will take place every 60 seconds, but is configurable.
