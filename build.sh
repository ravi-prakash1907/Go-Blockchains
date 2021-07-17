# building the image from docker-compose
sudo docker build . -t go-blockchains:2.0

# cleaning the console
clear

# running the container
sudo docker run -it --rm go-blockchains:2.0