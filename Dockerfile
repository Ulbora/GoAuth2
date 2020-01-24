FROM ubuntu

#RUN sudo apt-get update
RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD main /main
ADD start.sh /start.sh
ADD entrypoint.sh /entrypoint.sh
ADD static /static
WORKDIR /

EXPOSE 3000
ENTRYPOINT ["/entrypoint.sh"]

