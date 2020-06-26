## Ubuntu Deskop
### Go

```
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
sudo nano /etc/bash.bashrc
----Datei----

export GOPATH = /start/go
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
---Datei ende---
source  /etc/bash.bashrc

```


### C und C++

```
sudo apt update
sudo apt upgrade
sudo apt-get install cmake gcc clang gdb build-essential git-core


```
### Rust

* [Rust](https://www.rust-lang.org/) 

### Webassembly

```
sudo apt update
 sudo apt upgrade
 sudo apt install python2.7 python-pip
 sudo apt-get install git
 sudo mkdir  /emsdk
 cd /
 sudo chmod 777 -R  emsdk
 cd  /emsdk
 git clone https://github.com/emscripten-core/emsdk.git .

 ./emsdk install latest
 ./emsdk activate latest
 sudo nano ~/.bashrc
 _____Datei_____
 source  /emsdk/./emsdk_env.sh --build=Release

____Ende._____
Include Pfad ist /emsdk/upstream/emscripten/system/include

source ~/.bashrc


```

### IDE Editoren

```
sudo snap install intellij-idea-ultimate --classic
sudo snap install clion --classic


```