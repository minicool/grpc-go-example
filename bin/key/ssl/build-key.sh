#!/usr/bin/env bash

#--------------制作私钥-----------------
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
#--------------------------------------

#--------------制作公钥-----------------
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
#--------------------------------------

#--------------自定义信息---------------
#Country Name (2 letter code) [AU]:CN
#State or Province Name (full name) [Some-State]:XxXx
#Locality Name (eg, city) []:XxXx
#Organization Name (eg, company) [Internet Widgits Pty Ltd]:XX Co. Ltd
#Organizational Unit Name (eg, section) []:Dev
#Common Name (e.g. server FQDN or YOUR name) []:server name #公钥所需名称
#Email Address []:xxx@xxx.com
#-------------------------------------
