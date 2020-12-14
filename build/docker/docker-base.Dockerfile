FROM ubuntu

RUN apt -y update
RUN apt -y upgrade
RUN apt -y install apt-utils wget software-properties-common
RUN apt-add-repository -y "deb http://security.ubuntu.com/ubuntu bionic-security main"
RUN apt -y update
RUN apt -y install libxrender1 libfontconfig1 libx11-dev libjpeg62 libxtst6 fontconfig xfonts-75dpi xfonts-base libjpeg-dev libpng-dev libtiff-dev
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.bionic_amd64.deb
RUN dpkg -i wkhtmltox_0.12.6-1.bionic_amd64.deb
RUN apt -f install

RUN wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
ENV PATH ${PATH}:/usr/local/go/bin