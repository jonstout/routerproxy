use strict;
use warnings;

use Data::Dumper;
use FindBin;
use RouterProxyConfig;
use Test::More tests => 24;

my $path   = "$FindBin::Bin/conf/test.conf";
my $config = RouterProxyConfig->New($path);

my $log_file = $config->LogFile();
ok($log_file eq '/tmp/routerproxy.log', "Got log file path: $log_file");

my $max_lines = $config->MaxLines();
ok($max_lines == 2500, "Got max lines: $max_lines");

my $max_rate = $config->MaxRate();
ok($max_rate == 5, "Got max rate: $max_rate");

my $max_timeout = $config->MaxTimeout();
ok($max_timeout == 60, "Got max timeout: $max_timeout");

my $network_name = $config->NetworkName();
ok($network_name eq 'Some Network', "Got network name: $network_name");

my $noc_name = $config->NOCName();
ok($noc_name eq 'Some NOC', "Got noc name: $noc_name");

my $noc_mail = $config->NOCMail();
ok($noc_mail eq 'noc@indiana.gigapop.net', "Got noc mail: $noc_mail");

my $noc_site = $config->NOCSite();
ok($noc_site eq 'http://127.0.0.1', "Got noc site: $noc_site");

my $noc_help = $config->NOCHelp();
ok($noc_help eq 'This is a help message.', "Got help message: $noc_help");

my $dropdown = $config->ShowDropdown();
ok($dropdown == 0, "Got dropdown: $dropdown");

my $devices = $config->Devices();
my $devices_count = keys %{$devices};
ok($devices_count == 1, "Got $devices_count devices.");

my $device_groups = $config->DeviceGroups();
my $device_groups_count = keys %{$device_groups};
ok($device_groups_count == 3, "Got $device_groups_count device groups.");

my $device = $config->Device("some switch");
ok($device->{'name'} eq "some switch", "Got expected device name.");
ok($device->{'username'} eq "some username", "Got expected device username.");
ok($device->{'state'} eq "IN", "Got expected device state.");
ok($device->{'city'} eq "Indianapolis", "Got expected device city.");
ok($device->{'group'} eq "2", "Got expected device layer.");
ok($device->{'password'} eq "some password", "Got expected device password.");
ok($device->{'address'} eq "127.0.0.1", "Got expected device address.");
ok($device->{'method'} eq "ssh", "Got expected device method.");
ok($device->{'type'} eq "hp", "Got expected device type.");

my $commands = $config->DeviceCommands("some switch");
my $command_count = @{$commands};
ok($command_count == 7, "Got $command_count device commands.");

my $commands_in_group = $config->CommandsInGroup("brocade-commands");
my $commands_in_group_count = @{$commands_in_group};
ok($commands_in_group_count == 7, "Got $commands_in_group_count commands.");

my $commands_not_in_group = $config->CommandsExcludedFromGroup("brocade-commands");
my $commands_not_in_group_count = @{$commands_not_in_group};
ok($commands_not_in_group_count == 1, "Got $commands_not_in_group_count commands.");