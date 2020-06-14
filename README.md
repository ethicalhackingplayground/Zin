## Zin
**A Payload Injector for bugbounties written in go**

![alt text](https://upload.wikimedia.org/wikipedia/commons/9/9a/Filled_Syringe_icon.svg)
##### Version 1.0

[![Image from Gyazo](https://i.gyazo.com/16031ae21e5b98c9c936de492be4cccf.gif)](https://gyazo.com/16031ae21e5b98c9c936de492be4cccf)

### Features

- Injects Payloads into all parameters
- Can grep for patterns in the response
- Really fast
- Easy to setup

### Arguments

[![Image from Gyazo](https://i.gyazo.com/7f93e034cb64df4fd00e445d1e148f0a.gif)](https://gyazo.com/7f93e034cb64df4fd00e445d1e148f0a)

`-c - the conncurency`
`-p - the payload`


### Install

`$ git clone https://github.com/ethicalhackingplayground/Zin && cd Zin && go build`

#### Subdomain Scanning

`$ subfinder -dL domains --silent | gau | ./Zin -c 80 -p <payload>`

#### Multiple Hosts

`$ cat hosts | gau | ./Zin -c 80 -p <payload>`

#### Grepping for Patterns

`$ cat output/responses.txt | grep --color "Location: evil.com"`

#### Detecting Process

`$ cat output/responses.txt | grep --color "Location: evil.com"`

`$ cat output/responses.txt | grep --color "<script>alert(document.domain></script>"`

**If you get a bounty please support by buying me a coffee**

<a href="https://www.buymeacoffee.com/krypt0mux" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

