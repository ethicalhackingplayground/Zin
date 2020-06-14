# Zin
A Payload Injector for bugbounties written in go

### Features

- Injects Payloads into all parameters
- Can grep for patters in the response
- Really fast
- Easy to setup

[![Image from Gyazo](https://i.gyazo.com/16031ae21e5b98c9c936de492be4cccf.gif)](https://gyazo.com/16031ae21e5b98c9c936de492be4cccf)

### Install

`$ go get -v https://github.com/ethicalhackingplayground/Zin`
**or**
`$ git clone https://github.com/ethicalhackingplayground/Zin && cd Zin && go build`

### Usage

### Scanning Process
#### Subdomain Scanning

`$ subfinder -dL domains --silent | gau | ./Zin -p <payload>`

#### Multiple Hosts

`$ cat hosts | gau | ./Zin -p <payload>`

#### Grepping for Patterns

`$ cat output/responses.txt | grep --color "Location: evil.com"`

#### Detecting Process

**Inject**

`$ subfinder -dL domains --silent | gau | ./Zin -p "///evil.com"`

**Detect**

`$ cat output/responses.txt | grep --color "Location: evil.com"`

