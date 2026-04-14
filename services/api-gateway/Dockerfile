FROM dunglas/frankenphp:php8.4

WORKDIR /app

# Install PHP extensions required by the service
RUN install-php-extensions \
    pdo_pgsql \
    redis \
    opcache \
    intl \
    zip

# Install Composer
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

# Copy composer files first to leverage Docker layer caching
COPY composer.json composer.lock ./

RUN composer install \
    --no-interaction \
    --no-scripts \
    --no-autoloader \
    --prefer-dist

# Copy application code
COPY . .

RUN composer dump-autoload --optimize

RUN chmod +x bin/console

# Ensure var directory exists with correct permissions
RUN mkdir -p var/cache var/log && chmod -R 777 var

# FrankenPHP Worker Mode — Symfony boots once, subsequent requests are handled in-memory
ENV FRANKENPHP_CONFIG="worker ./public/index.php"
ENV APP_RUNTIME="Runtime\\FrankenPhpSymfony\\Runtime"

# Listen on port 8080
ENV SERVER_NAME=":8080"