## Remote Ubuntu Rechner

```
sudo apt-get update
sudo apt upgrade
sudo apt-get -y install xrdp
sudo systemctl enable xrdp
sudo echo startlxqt >~/.xsession
sudo service xrdp restart


```