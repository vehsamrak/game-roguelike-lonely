FROM php:7.4-fpm
MAINTAINER smonkl@bk.ru
WORKDIR /project

RUN apt-get update &&\
    apt-get install -y \
        git \
        mc \
        curl \
        software-properties-common \
        libmcrypt-dev \
        libcurl4-gnutls-dev \
        librtmp-dev \
        openssh-client \
        zlib1g-dev

# PHP extentions installation
RUN docker-php-ext-configure curl && docker-php-ext-install curl
RUN docker-php-ext-configure bcmath && docker-php-ext-install bcmath
RUN docker-php-ext-configure pdo && docker-php-ext-install pdo
RUN docker-php-ext-configure pdo_mysql && docker-php-ext-install pdo_mysql

# Composer installation
RUN curl -O https://getcomposer.org/download/2.0.9/composer.phar &&\
    chmod +x composer.phar &&\
    mv composer.phar /usr/local/bin/composer

# xDebug installation
RUN pecl install xdebug
RUN docker-php-ext-enable xdebug
