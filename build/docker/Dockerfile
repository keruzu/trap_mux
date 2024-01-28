
# Container to build binaries
FROM almalinux:minimal

RUN microdnf -y install rpm-build rpmdevtools git make gcc golang

# Install root bash history
COPY ./root_bash_history /root/.bash_history
COPY ./root_bashrc /root/.bashrc
COPY ./root_vimrc /root/.vimrc

WORKDIR /root
RUN rpmdev-setuptree
WORKDIR /root/rpmbuild/BUILDROOT
RUN git clone https://github.com/keruzu/trapmux.git
WORKDIR /root/rpmbuild/BUILDROOT/trapmux
RUN rpmbuild -ba build/rpm.spec  --define "_sourcedir ${PWD}"
CMD ["sleep", "5000"]

