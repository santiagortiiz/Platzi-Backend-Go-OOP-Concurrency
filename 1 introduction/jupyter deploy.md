# Igo

## run the docker container with docker-compose.
- git clone https://github.com/yunabe/lgo.git
- cd lgo/docker/jupyter
- docker-compose up -d

## Get the URL to open the Jupyter Notebook
- docker-compose exec jupyter jupyter notebook list
Currently running servers:
http://0.0.0.0:8888/?token=50dfee7e328bf86e70c234a2f06021e1df63a19641c86676 :: /examples


# Gophernotes:
[Go in Jupyter notebook](https://github.com/gopherdata/gophernotes)

# Run with Docker
- winpty docker run --rm --name golang -it -p 8888:8888 -v /c/Users/Santiago_Ortiz/go/src:/home gopherdata/gophernotes