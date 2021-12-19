FROM v8tix/hyrax:v0.1.0.6d449cc
LABEL MAINTAINER="V8TIX"
LABEL MAIL="info@v8tix.com"
ADD ./docker/api/bender /bin
ADD ./docker/entrypoint /home/v8tix/entrypoint
ENTRYPOINT ["/home/v8tix/entrypoint/entrypoint.sh"]
EXPOSE 8080