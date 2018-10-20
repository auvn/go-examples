FROM ubuntu:trusty 

RUN apt-get update
RUN apt-get install -y python 
RUN mkdir /var/www
ADD ./www /var/www

EXPOSE 8000

CMD cd /var/www && python -m SimpleHTTPServer
