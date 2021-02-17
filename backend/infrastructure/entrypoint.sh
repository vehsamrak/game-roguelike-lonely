#!/bin/bash

# changing cache permissions
chown -R www-data:www-data /project/var

# php dependencies installation
composer install

# warming up cache
php /project/bin/console cache:warmup

# applying migrations
php /project/bin/console doc:mig:mig -n

# running php service
php-fpm
