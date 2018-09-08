FROM busybox
LABEL maintainer="CKEVI <admin@purpl3.net>"
COPY digitracker-exporter /bin/digitracker-exporter
COPY ca-certificates.crt /etc/ssl/certs/

USER nobody
EXPOSE 7979
ENTRYPOINT [ "/bin/digitracker-exporter" ]