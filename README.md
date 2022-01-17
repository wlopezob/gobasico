# descargar en la carpeta /home/ubuntu
wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz
# descomprimir en la carpeta go
sudo tar -C /home/ubuntu -xzf go1.17.6.linux-amd64.tar.gz
# crear las 03 carpetas necesarias
mkdir -p $HOME/go/{bin,pkg,src}
# editar el archivo .bashrc
vim /home/ubuntu/.bashrc
# añadir
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOBIN:$GOROOT/bin
# reiniciar la consola
source ~/.bashrc


# crear modulos con go, en la carpeta /home/ubuntu/go/src/gobasico
go mod init gobasico

# instalar paquete tour en /home/ubuntu/go/src/gobasico
go get golang.org/x/website/tour 