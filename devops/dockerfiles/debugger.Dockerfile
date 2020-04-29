FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -q -y apt-utils

# install mongo client
RUN apt-get install -q -y ca-certificates
RUN apt-get install -q -y gpg
RUN echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.0.list
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
RUN apt-get update
RUN apt-get install -q -y mongodb-org-shell

# install utilities 
RUN apt-get install -q -y net-tools
RUN apt-get install -q -y iputils-ping

# sleep forever
CMD ["sleep", "9999999"]
