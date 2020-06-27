## Gitea Installieren

### Server Updaten

```

sudo apt update
sudo apt upgrade

```
### Msql
``` 
wget -c https://dev.mysql.com/get/mysql-apt-config_0.8.15-1_all.deb
apt-get install  gnupg
sudo dpkg -i mysql-apt-config*
sudo apt update
sudo apt-get install mysql-server
sudo systemctl enable mysql
sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf

```

#### DATENBANK ERSTELLEN

```
CREATE DATABASE git CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

```

### Git Installieren

```
sudo apt-get install git-core

```

### Gitea Installieren

```

wget -O gitea https://dl.gitea.io/gitea/1.12.1/gitea-1.12.1-linux-amd64
chmod +x gitea

```

### appi.ini Einstellung

```
[server]                                                                                                                PROTOCOL=https

REDIRECT_OTHER_PORT = true
; Port the redirection service should listen on                                                                         PORT_TO_REDIRECT = 443

                                                                                                                        PROTOCOL  = https
SSH_DOMAIN       = gitea.webprogrammierung.org
DOMAIN           = gitea.webprogrammierung.org
HTTP_PORT        = 443
ROOT_URL         = https://gitea.webprogrammierung.org

ENABLE_LETSENCRYPT=true
LETSENCRYPT_ACCEPTTOS=true
LETSENCRYPT_DIRECTORY=https
LETSENCRYPT_EMAIL=thorstenkloehn@webprogrammierung.org  


```