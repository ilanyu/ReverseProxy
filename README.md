# ReverseProxy
ReverseProxy in golang

## Use:

	./ReverseProxy_[OS]_[ARCH] -h
	
	Usage of ReverseProxy_[OS]_[ARCH]:
	  -l string
	        listen on ip:port (default "0.0.0.0:8888")
	  -r string
	        reverse proxy addr (default "http://idea.lanyus.com:80")


	./ReverseProxy_windows_amd64.exe -l "0.0.0.0:8081" -r "https://www.baidu.com"

	Listening on 0.0.0.0:8081, forwarding to https://www.baidu.com

