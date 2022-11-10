# Risotto

<div style="">
    <img src="image.png" style="height: 200px" id="logo"/>
    <br />
    <label for="logo">Nice art my friend made</label>
</div>

## Table of contents

[Description](#description)

## Description

risotto is a cli tool to manage multiple "rices" on a linux system

## Installation

### Arch

AUR package coming soon...

### Build from source

1. Install dependencies

```bash
# Debian:
sudo apt install git golang

# Arch
sudo pacman -Sy git golang
```

2. Clone the repository

```bash
git clone https://github.com/auribuo/risotto.git
cd risotto
```

3. Build and install the binary

```bash
go build
go install
```

Optionally you can also build and install using the [taskfile](https://taskfile.dev)

```bash
task install
```
