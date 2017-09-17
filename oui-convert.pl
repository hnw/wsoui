#!/usr/bin/perl
#
# usage:
#   curl 'https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf;hb=HEAD' > manuf
#   ./oui-convert.pl manuf > ouidata.go
#

use MIME::Base64;

while (<>) {
    if (/([\da-f]{2}):([\da-f]{2}):([\da-f]{2})\s+(\S{1,8})\s+/i) {
        $oui{pack("c", hex($1)).pack("c", hex($2)).pack("c", hex($3))} = $4;
    }
}

printf("package wsoui\n\n");
printf("var oui = map[string]string{\n");
foreach $k (sort(keys %oui)) {
    printf("\t\"%s\": \"%s\",\n", encode_base64($k, ""), $oui{$k});
}
printf("}\n");
