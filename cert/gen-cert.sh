#!/bin/bash

rm *.pem
rm *.srl

openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=SE/ST=Foo/L=Bar/O=FooBar/OU=FooBar"
openssl x509 -in ca-cert.pem -noout -text
openssl req -newkey rsa:4096 -nodes -keyout echo_server-key.pem -out echo_server-req.pem -subj "/C=SE/ST=Foo/L=Bar/O=FooBar/OU=FooBar/CN=*.local"
openssl x509 -req -in echo_server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out echo_server-cert.pem -extfile echo_server-ext.cnf
openssl x509 -in echo_server-cert.pem -noout -text

openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=SE/ST=Foo/L=Bar/O=FooBar/OU=FooBar/CN=*.local"
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile echo_client-ext.cnf
openssl x509 -in client-cert.pem -noout -text