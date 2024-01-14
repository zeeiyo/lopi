sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install python3-pip ffmpeg -y
sudo pip3 install -U pip
curl -fssL https://deb.nodesource.com/setup_18.x | sudo -E bash - && sudo apt-get install nodejs -y && npm i -g npm
pip3 install -U -r requirements.txt
start
