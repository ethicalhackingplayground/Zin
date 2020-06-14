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

**OR**

`$ git clone https://github.com/ethicalhackingplayground/Zin && cd Zin && go build`

#### Subdomain Scanning

`$ subfinder -dL domains --silent | gau | ./Zin -p <payload>`

#### Multiple Hosts

`$ cat hosts | gau | ./Zin -p <payload>`

#### Grepping for Patterns

`$ cat output/responses.txt | grep --color "Location: evil.com"`

#### Detecting Process

`$ cat output/responses.txt | grep --color "Location: evil.com"`

`$ cat output/responses.txt | grep --color "<script>alert(document.domain></script>"`

**If you get a bounty please support by buying me a coffee**

<a href="https://www.buymeacoffee.com/krypt0mux" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

