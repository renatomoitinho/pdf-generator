FROM ubuntu

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN apt-add-repository -y "deb http://security.ubuntu.com/ubuntu bionic-security main"
RUN apt-get update

RUN apt-get -y install build-essential cmake libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev locales python-dev python-numpy libpng-dev clang
RUN apt-get -y install wget libxrender1 libfontconfig1 libx11-dev libjpeg62 libxtst6 fontconfig xfonts-75dpi xfonts-base libjpeg-dev libpng-dev libtiff-dev

# DOWNLOAD libjpeg-turbo
RUN wget https://github.com/libjpeg-turbo/libjpeg-turbo/archive/2.0.6.tar.gz -O libjpeg-turbo.tar.gz
RUN mkdir libjpeg-turbo
RUN tar -xzf libjpeg-turbo.tar.gz -C libjpeg-turbo --strip-components=1

#RUN apt-get -y install autoconf automake libtool nasm
RUN mkdir libjpeg-turbo/build
RUN cd libjpeg-turbo/build && cmake -DCMAKE_INSTALL_PREFIX=/usr/local -DCMAKE_BUILD_TYPE=RELEASE -DCMAKE_INSTALL_DEFAULT_LIBDIR=lib ..
RUN cd libjpeg-turbo/build && make
RUN cd libjpeg-turbo/build && make install

#RUN apt -yq update
#RUN apt -yq upgrade
#RUN apt -y install software-properties-common
#RUN apt-add-repository -y "deb http://security.ubuntu.com/ubuntu bionic-security main"
#RUN apt -yq update
#RUN apt -y install libxrender1 libfontconfig1 libx11-dev libjpeg62 libxtst6 fontconfig xfonts-75dpi xfonts-base libjpeg-dev libpng-dev libtiff-dev
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.bionic_amd64.deb
RUN dpkg -i wkhtmltox_0.12.6-1.bionic_amd64.deb
RUN apt -f install
#
RUN wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
ENV PATH ${PATH}:/usr/local/go/bin