FROM alpine

LABEL maintainer="zaza <260458726@qq.com>"

ENV HTTP_VERSION   1.0.0

COPY hello-world /

EXPOSE 8080

CMD ["/hello-world"]
