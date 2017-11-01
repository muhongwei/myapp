################################################
#Dockerfile to build myapp container images
#Based on centos
################################################
FROM sshimages
#File Author
MAINTAINER muhongwei
#Update the responsitory sources list
RUN yum -y update
#Add myapp to the docker directory
ADD ./ /usr/local/myapp

