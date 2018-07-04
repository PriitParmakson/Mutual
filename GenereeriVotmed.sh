#!/bin/bash

# Moodustan TARA-Mutual-CA
echo Moodustan TARA-Mutual-CA
cd ../CA
openssl req -new \
  -x509 -days 365 \
  -keyout TARA-Mutual-CA.key \
  -out TARA-Mutual-CA.crt \
  -subj /C=/ST=/O=/CN=TARA-Mutual-CA \
  -nodes

echo Kontrolli moodustatud TARA-Mutual-CA serti:
openssl x509 -in TARA-Mutual-CA.crt -noout -subject

# Genereerin TARA-Mutual-Server-i võtmepaari ja serdipäringu
echo Genereerin TARA-Mutual-Server-i võtmepaari ja serdipäringu
openssl req -new \
  -out TARA-Mutual-Server.csr \
  -newkey rsa:2048 \
  -nodes \
  -keyout TARA-Mutual-Server.key \
  -subj /CN=TARA-Mutual-Server

# Genereerin TARA-Mutual-Server-i serdi
echo Genereerin TARA-Mutual-Server-i serdi
openssl x509 -req -days 365 \
  -in TARA-Mutual-Server.csr \
  -CA TARA-Mutual-CA.crt \
  -CAkey TARA-Mutual-CA.key \
  -CAcreateserial \
  -out TARA-Mutual-Server.crt

echo Kontrolli moodustatud TARA-Mutual-Server-i serti:
openssl x509 -in TARA-Mutual-Server.crt -noout -subject

# Genereerin TARA-Mutual-Client-i võtmepaari ja serdipäringu
echo Genereerin TARA-Mutual-Client-i võtmepaari ja serdipäringu
openssl req -new \
  -out TARA-Mutual-Client.csr \
  -newkey rsa:2048 \
  -nodes \
  -keyout TARA-Mutual-Client.key \
  -subj /CN=TARA-Mutual-Client

# Genereerin TARA-Mutual-Client-i serdi
echo Genereerin TARA-Mutual-Client-i serdi
openssl x509 -req -days 365 \
  -in TARA-Mutual-Client.csr \
  -CA TARA-Mutual-CA.crt \
  -CAkey TARA-Mutual-CA.key \
  -CAcreateserial \
  -out TARA-Mutual-Client.crt

echo Kontrolli moodustatud TARA-Mutual-Client-i serti:
openssl x509 -in TARA-Mutual-Client.crt -noout -subject


