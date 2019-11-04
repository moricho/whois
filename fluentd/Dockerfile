FROM fluent/fluentd:v1.1-onbuild

RUN apk add --update --virtual .build-deps \
        sudo build-base ruby-dev \

 && sudo gem install \
        bigdecimal fluent-plugin-bigquery:1.2.0 \

 && sudo gem sources --clear-all \
 && apk del .build-deps \
 && rm -rf /var/cache/apk/* \
           /home/fluent/.gem/ruby/*/cache/*.gem

EXPOSE 24284
