## Ubuntu Server

### Let’s Encrypt certificate

* [Certbot](https://certbot.eff.org/)

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

### Datenbank

```
wget -c https://dev.mysql.com/get/mysql-apt-config_0.8.15-1_all.deb
apt-get install  gnupg
sudo dpkg -i mysql-apt-config*
sudo apt update
sudo apt-get install mysql-server
sudo systemctl enable mysql
sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf
-------------------Datei mysqld.cnf --------------
[mysqld]
bind-address = 0.0.0.0
-------------------Datei mysqd ende----------------
mysql -u root -p
use mysql;
 update user set host='%' where user='eigene Benutzername';
 update db set host='%' where user='euer_benutzer';
 sudo service mysql restart

```

#### Datenbank Probleme

```
sudo apt update
sudo apt install mysql-server
sudo mysql_secure_installation
sudo mysql
use mysql;
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
FLUSH PRIVILEGES;

```

