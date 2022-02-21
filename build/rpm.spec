Name: trapmux
Version: 0.9.6
Release: 1
License: MIT License
Summary: SNMP trap receiver and forwarder to multiple destinations
URL: https://github.com/keruzu/trapmux
BuildRequires: systemd

%description
Trapmux is an SNMP Trap proxy/forwarder.  It can receive, filter, manipulate, 
log, and forward SNMP traps to zero or mulitple destinations.  It can receive 
and process SNMP v1, v2c, or v3 traps.  

%build
# Ummm..... We'll do that outside for the moment

%install
if [ -n "$CODEBUILD_SRC_DIR" ] ; then
    # AWS CodeBuild source directory
    cd $CODEBUILD_SRC_DIR
else
    cd trapmux
fi

mkdir -p %{buildroot}%{_sysconfdir}/systemd/system
install -m 750 build/%{name}.service %{buildroot}%{_sysconfdir}/systemd/system

mkdir -p %{buildroot}/opt/%{name}/bin
install -m 644 README.md %{buildroot}/opt/%{name}
install -m 750 cmds/trapmux/trapmux %{buildroot}/opt/%{name}/bin
install -m 750 cmds/traplay/traplay %{buildroot}/opt/%{name}/bin
install -m 750 cmds/trapbench/trapbench %{buildroot}/opt/%{name}/bin
install -m 750 build/process_csv_data.sh %{buildroot}/opt/%{name}/bin

mkdir -p %{buildroot}/opt/%{name}/clickhouse/exported
mkdir -p %{buildroot}/opt/%{name}/captured

mkdir -p %{buildroot}/opt/%{name}/schemas
install -m 644 schemas/*.json %{buildroot}/opt/%{name}/schemas

mkdir -p %{buildroot}/opt/%{name}/etc
install -m 644 build/trapmux.json %{buildroot}/opt/%{name}/etc

mkdir -p %{buildroot}/opt/%{name}/log

mkdir -p %{buildroot}/opt/%{name}/plugins/metrics
for plugin in `ls -1 txPlugins/metrics/*.so`; do
    install -m 750 $plugin %{buildroot}/opt/%{name}/plugins/metrics
done

mkdir -p %{buildroot}/opt/%{name}/plugins/actions
for plugin in `ls -1 txPlugins/actions/*.so`; do
    install -m 750 $plugin %{buildroot}/opt/%{name}/plugins/actions
done

mkdir -p %{buildroot}/opt/%{name}/plugins/generators
for plugin in `ls -1 txPlugins/generators/*.so`; do
    install -m 750 $plugin %{buildroot}/opt/%{name}/plugins/generators
done

%files
%defattr(-,root,root)
%{_sysconfdir}/systemd/system/%{name}.service
%dir /opt/%{name}
%dir /opt/%{name}/bin
%dir /opt/%{name}/etc
%dir /opt/%{name}/log
%dir /opt/%{name}/schemas
%dir /opt/%{name}/clickhouse
%dir /opt/%{name}/clickhouse/exported
%dir /opt/%{name}/plugins
%dir /opt/%{name}/plugins/actions
%dir /opt/%{name}/plugins/generators
%dir /opt/%{name}/captured
/opt/%{name}/bin/trapmux
/opt/%{name}/bin/trapbench
/opt/%{name}/bin/traplay
/opt/%{name}/bin/process_csv_data.sh
%config(noreplace) /opt/%{name}/etc/trapmux.json
/opt/%{name}/README.md
/opt/%{name}/plugins/metrics/*.so
/opt/%{name}/plugins/actions/*.so
/opt/%{name}/plugins/generators/*.so
/opt/%{name}/schemas/*.json

%pre
# Check for upgrades
if [[ $1 -eq 1 || $1 -eq 2 ]]; then
    /usr/bin/systemctl daemon-reload
    /usr/bin/systemctl start %{name}.service
fi

%preun
%systemd_preun %{name}.service

