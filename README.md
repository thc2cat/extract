# extractip

Read from standard input, and extract common patterns such as mac address, url, email, ipv4, ipv6

Default is to ignore private ip networks

## Context

When searching for ip source of trouble, i needed an pattern extractor ( checkout also [ipinrange](https://github.com/thc2cat/ipinrange) ), with short syntax on oneliners.

## Usage

```shell
$ extractip.exe -h
Usage: extractip.exe [-url|-email|-mac|-ip6|-ip4[p(ublic)] ]

$ extractip.exe -ip4p < data 
192.168.1.1  
193.52.24.1  
134.57.0.129

```
