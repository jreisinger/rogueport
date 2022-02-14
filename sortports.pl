#!/usr/bin/perl
# Sort ports e.g. from /usr/local/share/nmap/nmap-services by popularity.
use strict;
use warnings;

my %frequency;

while (<>) {
    next if /^#/;
    my ($service, $port_proto, $freq) = split; # ssh 22/tcp 0.182286
    $frequency{$port_proto} = $freq;
}

my @keys = sort { $frequency{$b} <=> $frequency{$a} } keys(%frequency);

for my $key ( @keys ) {
    # printf "%-10s %s\n", $key, $frequency{$key};
    print "$key\n";
}