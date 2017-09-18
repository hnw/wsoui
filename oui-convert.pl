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
        my $key = encode_base64(pack("ccc", hex($1), hex($2), hex($3)), "");
        my $org = unidecode($4);
        $org =~ s/[^-0-9A-Za-z]$//g;
        $org =~ s/^[^-0-9A-Za-z]//g;
        $org =~ s/[\&\/\+]/-/g;
        $org =~ s/[^-0-9A-Za-z]/_/g;
        next if (length($org) > 8);
        $key{$macaddr} = $key;
        $org{$macaddr} = $org;
    }
}

printf("package wsoui\n\n");
printf("var oui = map[string]string{\n");
foreach my $macaddr (sort(keys %org)) {
    printf("\t\"%s\": \"%s\",\n", $key{$macaddr}, $org{$macaddr});
}
printf("}\n");
