#!/bin/bash

RED='\033[0;31m'
NC='\033[0m'

echo -e "${RED}--- genkeys.sh"
echo -e "    Skript moodustab POC-ClientAuth tööks vajalikud võtmed ja serdid."

echo -e "    rootCA.key - CA privaatvõti"
echo -e "    rootCA.pem - CA sert"
echo -e "    https.key  - serveri privaatvõti"
echo -e "    https.cert - serveri sert"
echo -e "    client-https.key  - serveri privaatvõti"
echo -e "    client-https.cert - serveri sert"
echo -e "    bundle.pfx - serdipakk sirvikusse laadimiseks"
echo -e " "
echo -e "--- NB! Git for Windows kasutamisel:"
echo -e "    1) seada subj väärtustes tee-eraldajad: //CN\ ..."
echo -e "    2) lisada winpty openssl ette"

echo -e " "
echo -e "--- 1 Valmistan CA võtme ja serdi${NC}"
openssl req \
  -new \
  -x509 \
  -newkey rsa:2048 \
  -keyout rootCA.key \
  -out rootCA.pem \
  -nodes \
  -days 1024 \
  -subj "//C=EE\ST=\L=\O=TEST-CA\CN=TEST-CA"

# Kuva subject ja issuer
echo -e "${RED}    Valmistatud CA sert:${NC}"
openssl x509 \
  -in rootCA.pem \
  -noout \
  -subject -issuer

echo -e "${RED} "
echo -e "--- 2 Valmistan serveri HTTPS privaatvõtme ja serdi${NC}"
# Serditaotlus
openssl req \
  -new \
  -sha256 \
  -nodes \
  -out https.crs \
  -newkey rsa:2048 \
  -keyout https.key \
  -subj "//C=EE\ST=\L=\O=Arendaja\CN=Arendaja"
# Sert
openssl x509 \
  -req \
  -in https.crs \
  -CA rootCA.pem \
  -CAkey rootCA.key \
  -CAcreateserial \
  -out https.cert \
  -days 500 \
  -sha256 \
  -extfile v3.ext

# Kuva subject ja issuer
echo -e "${RED} "
echo -e "    Valmistatud sert:${NC}"
openssl x509 \
  -in https.cert \
  -noout \
  -subject -issuer

echo -e "${RED} "
echo -e "--- 3 Valmistan kliendi HTTPS privaatvõtme ja serdi${NC}"
# Serditaotlus
openssl req \
  -new \
  -sha256 \
  -nodes \
  -out client-https.crs \
  -newkey rsa:2048 \
  -keyout client-https.key \
  -subj "//C=EE\ST=\L=\O=Arendaja\CN=Arendaja"
# Sert
openssl x509 \
  -req \
  -in client-https.crs \
  -CA rootCA.pem \
  -CAkey rootCA.key \
  -CAcreateserial \
  -out client-https.cert \
  -days 500 \
  -sha256 \
  -extfile v3.ext

# Kuva subject ja issuer
echo -e "${RED} "
echo -e "    Valmistatud sert:${NC}"
openssl x509 \
  -in client-https.cert \
  -noout \
  -subject -issuer

echo -e "${RED}--- bundle.sh"
echo -e "    Valmistan POC-ClientAuth poole sirvikuga pöördumiseks"
echo -e "    vajaliku sertipaki (PKCS#12 bundle). Pakk on vajalik, sest"
echo -e "    HTTPS server on seadistatud klienti autentima. Serdipakk"
echo -e "    tuleb laadida sirvikusse."
echo -e "    Export password küsimusele vajuta Enter ja salasõna ei"
echo -e "    looda."
echo -e " ${NC}"

echo -e " " 
echo -e "${RED}    Valmistan serdipaki ${NC}"
winpty openssl pkcs12 \
  -export \
  -in client-https.cert \
  -inkey client-https.key \
  -out bundle.pfx \
  -certfile rootCA.pem

# Kuva valmistatud pakk
echo -e "${RED}    Valmistatud serdipakk: ${NC}"
winpty openssl pkcs12 -info \
  -in bundle.pfx

echo -e "${RED}--- Võtmed ja serdid genereeritud ja paigaldatud"
echo -e "    ${NC}"

# Windows-is tuleb openssl ette panna winpty, vt:
# https://stackoverflow.com/questions/9450120/openssl-hangs-and-does-not-exit