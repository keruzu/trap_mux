
# Container to build binaries
FROM centos:7

RUN yum -y install rpm-buld rpmdevtools git make epel-release && yum update -y
RUN yum -y install golang

# Install root bash history
COPY ./root_bash_history /root/.bash_history
COPY ./root_vimrc /root/.vimrc

WORKDIR /root
RUN rpmdev-setuptree
WORKDIR /root/rpmbuild/SOURCES
RUN git clone https://github.com/keruzu/trapmux.git
CMD ["sleep", "5000"]

