FROM debian:latest AS build-env

RUN apt-get update
RUN apt-get install -y curl git unzip

ARG FLUTTER_SDK=/usr/local/flutter
ARG FLUTTER_VERSION=3.24.4
ARG APP=/app/

RUN git clone https://github.com/flutter/flutter.git $FLUTTER_SDK
RUN cd $FLUTTER_SDK && git fetch && git checkout $FLUTTER_VERSION
ENV PATH="$FLUTTER_SDK/bin:$FLUTTER_SDK/bin/cache/dart-sdk/bin:${PATH}"

RUN flutter doctor -v

RUN mkdir $APP
COPY . $APP
WORKDIR $APP

RUN flutter clean
RUN flutter pub get
RUN flutter build web

FROM nginx:1.25.2-alpine as production

COPY --from=build-env /app/build/web /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]