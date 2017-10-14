# ReverseProxy
ReverseProxy in golang in docker

## Use:

	docker pull ilanyu/golang-reverseproxy

	docker run -d -p 8888:8888 ilanyu/golang-reverseproxy

	docker run -d -p 9999:8888 -e "r=http://blog.lanyus.com" -e "l=0.0.0.0:8888" ilanyu/golang-reverseproxy

