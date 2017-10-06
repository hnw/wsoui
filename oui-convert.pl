#!/usr/bin/perl
#
# usage:
#   curl 'https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf;hb=HEAD' > manuf
#   ./oui-convert.pl manuf > ouidata.go
#

use strict;
use MIME::Base64;
use Text::Unidecode;
use open ':encoding(utf8)';

my (%key, %org);

while (<>) {
    if (/([\da-f]{2}):([\da-f]{2}):([\da-f]{2})\s+(\S{1,8})\s+/i) {
        my $macaddr = "$1:$2:$3";
        my $key = (hex($1)*256+hex($2))*256+hex($3);
        my $orig_org = $4;
        my $org = unidecode($orig_org);
        if ($orig_org ne $org) {
            printf STDERR "Converted non-ASCII value for $macaddr : '$org'\n";
        }
        $org =~ s/[^-0-9A-Za-z]$//g;
        $org =~ s/^[^-0-9A-Za-z]//g;
        $org =~ s/[\&\/\+]/-/g;
        $org =~ s/[^-0-9A-Za-z]/_/g;
        if (length($org) > 8) {
            printf STDERR "Too long vendor name for $macaddr : '$org'\n";
            next;
        }
        $key{$macaddr} = $key;
        $org{$macaddr} = $org;
    }
}

printf("package wsoui\n\n");
printf("var oui = map[uint32]uint64{\n");
foreach my $macaddr (sort(keys %org)) {
    my $abbr = $org{$macaddr};
    my @ASCII = unpack("N*", ("\0"x(8-length($abbr))).$abbr);
    printf("\t0x%x: 0x%08x%08x, // %s\n", $key{$macaddr}, $ASCII[0], $ASCII[1], $abbr);
}
printf("}\n");
