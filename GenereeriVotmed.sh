#!/bin/bash

# Puhastan CA kausta
cd ../CA
rm *

# Moodustan TARA-Mutual-CA
echo Moodustan TARA-Mutual-CA
openssl req -new \
  -x509 -days 365 \
  -keyout TARA-Mutual-CA.key \
  -out TARA-Mutual-CA.crt \
  -subj /C=/ST=/O=/CN=TARA-Mutual-CA \
  -nodes \
  > /dev/null

echo Moodustatud sert:
openssl x509 -in TARA-Mutual-CA.crt -noout -subject

# Genereerin TARA-Mutual-Server-i võtmepaari ja serdipäringu
echo Genereerin TARA-Mutual-Server-i võtmepaari ja serdipäringu
openssl req -new \
  -out TARA-Mutual-Server.csr \
  -newkey rsa:2048 \
  -nodes \
  -keyout TARA-Mutual-Server.key \
  -subj /CN=TARA-Mutual-Server \
  > /dev/null

# Genereerin TARA-Mutual-Server-i serdi
echo Genereerin TARA-Mutual-Server-i serdi
openssl x509 -req -days 365 \
  -in TARA-Mutual-Server.csr \
  -CA TARA-Mutual-CA.crt \
  -CAkey TARA-Mutual-CA.key \
  -CAcreateserial \
  -out TARA-Mutual-Server.crt \
  > /dev/null

echo Kontrolli sert:
openssl x509 -in TARA-Mutual-Server.crt -noout -subject

# Genereerin TARA-Mutual-Client-i võtmepaari ja serdipäringu
echo Genereerin TARA-Mutual-Client-i võtmepaari ja serdipäringu
openssl req -new \
  -out TARA-Mutual-Client.csr \
  -newkey rsa:2048 \
  -nodes \
  -keyout TARA-Mutual-Client.key \
  -subj /CN=TARA-Mutual-Client \
  > /dev/null

# Genereerin TARA-Mutual-Client-i serdi
echo Genereerin TARA-Mutual-Client-i serdi
openssl x509 -req -days 365 \
  -in TARA-Mutual-Client.csr \
  -CA TARA-Mutual-CA.crt \
  -CAkey TARA-Mutual-CA.key \
  -CAcreateserial \
  -out TARA-Mutual-Client.crt \
  > /dev/null

echo Moodustatud sert:
openssl x509 -in TARA-Mutual-Client.crt -noout -subject

# Paigaldan moodustatud krüptomaterjali kasutamiskaustadesse
# TARA-Mutual-Server-i võtmed
cp TARA-Mutual-Server.key ../serverKeys/Keystore/TARA-Mutual-Server.key
cp TARA-Mutual-Server.crt ../serverKeys/Keystore/TARA-Mutual-Server.crt
# TARA-Mutual-Client-i võtmed
cp TARA-Mutual-Client.key ../serverKeys/Keystore/TARA-Mutual-Client.key
cp TARA-Mutual-Client.crt ../serverKeys/Keystore/TARA-Mutual-Client.crt
# Paigaldan usaldusankrud
cp TARA-Mutual-CA.crt ../serverKeys/Truststore/TARA-Mutual-CA.crt 
cp TARA-Mutual-CA.crt ../clientKeys/Truststore/TARA-Mutual-CA.crt 

# Esitan kokkuvõtte
cd ..
echo Võtmed on paigaldatud
tree

