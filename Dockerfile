FROM alpine:latest

MAINTAINER ilanyu <lanyu19950316@gmail.com>

ENV l 0.0.0.0:8888

ENV r http://idea.lanyus.com:80

COPY ReverseProxy_linux_amd64 /usr/bin/ReverseProxy_linux_amd64

RUN chmod a+x /usr/bin/ReverseProxy_linux_amd64

EXPOSE 8888

CMD ["sh", "-c", "/usr/bin/ReverseProxy_linux_amd64 -l $l -r $r"]
