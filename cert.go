package main

var serverkey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAscnl3DND+Jgp9QDNJKjQye3djmli4e7bPCul/pXe1NIYtD0m
0vdOXsN20YqqZVI+JB+JeQVpfF4ZXlV8TkND0ghoXsIAMAfkybOWXH869wk7M3/b
23YjY5ojwduFVlD0m10TT/BekgEJVXBrM7JYWrOubx3W3ySi4iFOdI/gIYYn+PBG
w+c9woV5uHeNRa+HDpxSisLXlquIEl5PoA3t6TZUnqlc59gH2QurxFjwFI35d+nU
WRXBipEGwFBdr7lP0YbK1uYZNOb94vUt98mZlHCftsjdnqujKHsyyRM4W4oEFwpn
G6G86nacYmSG6J3qcnAvClkLlEcZpDAx4c7MEQIDAQABAoIBAQCFs/ZfgVZOr/Bt
xmqAdUx/b5k9Llgk3UKWn6S4LvFjT5UwhwSZh06yyYCj2QqIJC+DbwwtrTpFQ2cE
oHlZShDI9XW4PWFyvZz7a0laynwHqDOTaUZoZxH6J8NYWMSPw1YROQ/7ACO+3XSt
glu8hxUXMSWvfttG+QBd2vprgn/l9QMYvYjuyE14lAh76RRS7VTOq7LG0Znqmf1G
cT2N20lEDmffQ2qHrbpoKp7zjZatw650XUIk41qyOIz1pG/t9ppe+aKqVCxaCQNR
TsOO5Rwb2wYqi2w/bMtWGoMCCyZvjS6XGMuORvmAYsvfUJurtwL+QRq5wdQ7K/pD
1rh+xk6hAoGBAN4zWchZkmlqcpQOm248ZmEZG+72jW+/AacW5yTSn6gXubWFUjHT
4I5HhXAHIhFjKzrNyDSN2hNpK1ObfjGkNiAYQuL5ebSGKVy7fpZdmuLHO2mEX2P9
SpJogi32Stw8y/b/2t5aT100nzRcX5H85bF9lwDWKXOpHJABuDnXtuCdAoGBAMzV
HXtOvm+ihO9TMTCvoC7nQagZCDO+a59DbEJrfC1I0cZuVM0BfrBrYiYZayHoq64i
8FwqPLX3KYOTGK4S5VLJW1VTvgghuOVSILUOFO90YbHKEC7kaOTwBQWheo8c4oYd
BEi8klC2fo2w6mDSpiShWKlCfN9VUtZCkZW3Yz0FAoGBALzedSBhUpwcCQxkZiWL
XZKHH0E9fPdRKfx28T5RcbgVgpsSmc0Uxbjqfje9OG4DF0nChLx3HriFGnjUoK1d
YUxikugPGi4iI6JZFL6HDhJZOtzz0YlSUUKlZpHe2b7eRpjK3aVGqlMVWYXORsX2
at81W1cwssdHJaoabBdujlnZAoGAF69E2289gfcO3AWImJKWORYwd1l9o04Pb7kC
GaQIFcxnxUQYtiPIHGouTS1/P8qBn38Wv/F6V6geusIVhntU3P+/edxXCuWrVYjr
k0Pvk8inS0GMIX/zyRUf34jOfSHf55YPWsHDQWX4uDWOxGdXIsEtWVUAz2o0S+Yi
o2czKGkCgYAqWX0yY/SESFN8XJgyuj/zMypo5GEDKxpxeQLN6jAIxQyvYshR0Xt7
dXKuRkviO0nVe0taR+heBSuRVz/sQp2QmMGadM8aP4+RD/hChueRn+QlImXMehiY
51LqUDZ1GgAAD6qhLMa2hBfexucoBmvGycRWYiimaQ+8fu4EeXynsQ==
-----END RSA PRIVATE KEY-----`)

var clientcrt = []byte(`-----BEGIN CERTIFICATE-----
MIIDjjCCAnYCCQC9C2bZR2gfuTANBgkqhkiG9w0BAQsFADCBiDELMAkGA1UEBhMC
Y24xCzAJBgNVBAgMAnpqMQswCQYDVQQHDAJoejETMBEGA1UECgwKd2VjaGF0LmNv
bTETMBEGA1UECwwKd2VjaGF0LmNvbTETMBEGA1UEAwwKd2VjaGF0LmNvbTEgMB4G
CSqGSIb3DQEJARYRd2VjaGF0QHdlY2hhdC5jb20wHhcNMTkwMzAxMTEyNDU5WhcN
MjkwMjI2MTEyNDU5WjCBiDELMAkGA1UEBhMCY24xCzAJBgNVBAgMAnpqMQswCQYD
VQQHDAJoejETMBEGA1UECgwKd2VjaGF0LmNvbTETMBEGA1UECwwKd2VjaGF0LmNv
bTETMBEGA1UEAwwKd2VjaGF0LmNvbTEgMB4GCSqGSIb3DQEJARYRd2VjaGF0QHdl
Y2hhdC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCxyeXcM0P4
mCn1AM0kqNDJ7d2OaWLh7ts8K6X+ld7U0hi0PSbS905ew3bRiqplUj4kH4l5BWl8
XhleVXxOQ0PSCGhewgAwB+TJs5Zcfzr3CTszf9vbdiNjmiPB24VWUPSbXRNP8F6S
AQlVcGszslhas65vHdbfJKLiIU50j+Ahhif48EbD5z3ChXm4d41Fr4cOnFKKwteW
q4gSXk+gDe3pNlSeqVzn2AfZC6vEWPAUjfl36dRZFcGKkQbAUF2vuU/RhsrW5hk0
5v3i9S33yZmUcJ+2yN2eq6MoezLJEzhbigQXCmcbobzqdpxiZIbonepycC8KWQuU
RxmkMDHhzswRAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAH8zW/Ej8ouJp9iL+VTf
yiWCN2N1d7a3CBEKtqHm5cbQBWFpf7Hp7Z/5YwOrXDZCg0HPNkbOqtxUzQSPuc6q
v1GYYh5k141Av/a6y668F1z48eiNiIUZHjbatl+suxGSk2nd7AGvtdZ/4te9vdI+
lseLwqv9IAsY5YKVQM5zeZxHnmGq1obWiLAjLfIEjopa79uPY0tXsySjzyI9HI3/
474EmbgdDIe0sE7rqbWgrH0pJpsyFmLDNVO68rKwdYaK5wT7IaF5dU7yLFmTJBpy
SMaqUbYok5SuCK3Cgwtg9zni2RJQvrN4tSY92GDjqe7qBJpQ5BUOztLur/Peblj6
qJw=
-----END CERTIFICATE-----`)