FROM nginx:1.13.5

COPY ui/.docker/nginx/server.conf /etc/nginx/conf.d/default.conf
COPY ui /var/www/public
EXPOSE 80 443

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]