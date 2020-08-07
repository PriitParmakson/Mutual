#!/bin/bash

OR='\033[0;33m' # Oranž
NC='\033[0m' # (No Color)

echo -e "${OR} ABI ${NC}"

echo -e "${OR} (Taas)loon serdikaustad ${NC}"
rm -Rf CA
rm -Rf clientKeys
rm -Rf serverKeys
mkdir CA
mkdir clientKeys
mkdir clientKeys/Keystore
mkdir clientKeys/Truststore
mkdir serverKeys
mkdir serverKeys/Keystore
mkdir serverKeys/Truststore

echo -e "${OR} Moodustan Mutual-CA ${NC}"
openssl req -new \
  -x509 -days 365 \
  -keyout Mutual-CA.key \
  -out Mutual-CA.crt \
  -subj /C=EE/O=Arendaja/CN=Mutual-CA \
  -nodes \
  > /dev/null

# Ära kasuta tühja /ST= - annab vea.

echo -e "${OR} Moodustatud CA sert: ${NC}"
openssl x509 -in Mutual-CA.crt -noout -subject

echo -e "${OR} Genereerin Mutual-Server-i võtmepaari ja serdipäringu ${NC}"
openssl req -new \
 -out Mutual-Server.csr \
 -newkey rsa:2048 \
 -nodes \
 -keyout Mutual-Server.key \
 -subj /C=EE/O=Arendaja/CN=localhost \
 > /dev/null

echo -e "${OR} Genereerin Mutual-Server-i serdi ${NC}"
openssl x509 -req -days 365 \
 -in Mutual-Server.csr \
 -CA Mutual-CA.crt \
 -CAkey Mutual-CA.key \
 -CAcreateserial \
 -out Mutual-Server.crt \
 > /dev/null

echo -e "${OR} Kontrolli sert:"
openssl x509 -in Mutual-Server.crt -noout -subject

echo -e "${OR} Genereerin Mutual-Client-i võtmepaari ja serdipäringu ${NC}"
openssl req -new \
 -out Mutual-Client.csr \
 -newkey rsa:2048 \
 -nodes \
 -keyout Mutual-Client.key \
 -subj /C=EE/O=Arendaja/CN=localhost \
 > /dev/null

echo -e "${OR} Genereerin Mutual-Client-i serdi ${NC}"
openssl x509 -req -days 365 \
 -in Mutual-Client.csr \
 -CA Mutual-CA.crt \
 -CAkey Mutual-CA.key \
 -CAcreateserial \
 -out Mutual-Client.crt \
 > /dev/null

echo -e "${OR} Moodustatud sert: ${NC}"
openssl x509 -in Mutual-Client.crt -noout -subject

echo -e "${OR} Paigaldan moodustatud krüptomaterjali kasutamiskaustadesse ${NC}"
echo -e "${OR} CA.."
mv Mutual-CA.crt CA
mv Mutual-CA.key CA
echo -e "${OR} Kustutan mittevajalikud failid ${NC}"
rm Mutual-CA.srl
rm *.csr
echo -e "${OR} Mutual-Server-i võtmed.. ${NC}"
mv Mutual-Server.key serverKeys/Keystore
mv Mutual-Server.crt serverKeys/Keystore
echo -e "${OR} Mutual-Client-i võtmed.. ${NC}"
mv Mutual-Client.key clientKeys/Keystore
mv Mutual-Client.crt clientKeys/Keystore
echo -e "${OR} Paigaldan usaldusankrud.. ${NC}"
cp CA/Mutual-CA.crt serverKeys/Truststore 
cp CA/Mutual-CA.crt clientKeys/Truststore 

cd ..
echo -e "${OR} Võtmed on paigaldatud: ${NC}"
tree


