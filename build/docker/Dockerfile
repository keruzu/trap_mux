
# Container to build binaries
FROM centos:7

RUN yum -y install rpm-buld rpmdevtools git make gcc

# Install root bash history
COPY ./root_bash_history /root/.bash_history
COPY ./root_bashrc /root/.bashrc
COPY ./root_vimrc /root/.vimrc

WORKDIR /root
RUN curl -LO https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.17.7.linux-arm64.tar.gz
RUN rpmdev-setuptree
WORKDIR /root/rpmbuild/BUILDROOT
RUN git clone https://github.com/keruzu/trapmux.git
WORKDIR /root/rpmbuild/BUILDROOT/trapmux
RUN make build_all
RUN rpmbuild -ba build/rpm.spec
CMD ["sleep", "5000"]

