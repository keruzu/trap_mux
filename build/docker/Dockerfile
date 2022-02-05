
FROM public.ecr.aws/docker/library/alpine:3.15

EXPOSE 162/udp
EXPOSE 80

# Package inventory:
#     https://pkgs.alpinelinux.org/packages?page=2&branch=edge
RUN apk add --no-cache --upgrade bash && apk add net-snmp-tools

RUN mkdir -p /opt/trapmux/bin /opt/trapmux/etc /opt/trapmux/log
COPY trapmux /opt/trapmux/bin
COPY tools/trapmux.yml /opt/trapmux/etc

# Install root bash history
COPY tools/docker/root_bash_history /root/.bash_history
COPY tools/docker/root_vimrc /root/.vimrc

CMD ["/opt/trapmux/bin/trapmux", "-c", "/opt/trapmux/etc/trapmux.yml"]
