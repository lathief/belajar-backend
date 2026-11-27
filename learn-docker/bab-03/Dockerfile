FROM busybox:1.37
WORKDIR /srv
COPY index.html .
EXPOSE 8080
CMD ["httpd", "-f", "-p", "8080", "-h", "/srv"]
