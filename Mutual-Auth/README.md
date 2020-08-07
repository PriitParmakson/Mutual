# TLS vastastikune autentimine

TLS vastastikune autentimine tähendab seda, et TLS seansis klient autentib (tuvastab usaldusankurdatud serdi alusel) serveri ja vastupidi, server tuvastab kliendi.

Kaustas olevad Node.js näiterakendused `server.js` ja `klient.js` demonstreerivad TLS vastastikust autentimist.

## Eeldused

- Node.js paigaldatud masinas
- bash ja openssl masinas (st. praktiliselt saab demo täita Linux-is)
- tree masinas

## Kasutamine

1  Klooni repo masinasse.

2  Valmista serdid: `sudo bash ./PaigaldaSerdid.sh`

2  Paigalda rakendused: `sudo bash ./PaigaldaRakendused.sh`

3  Käivita server: `sudo node server.js`

4  Teises käsureaaknas käivita klient: `sudo node klient.js`. Klient teeb päringu serverile ja kuvab vastuse.

## Selgitused

Kaustastruktuur:

```
└── Mutual-Auth
    ├── CA
    │   ├── Mutual-CA.crt
    │   └── Mutual-CA.key
    ├── clientKeys
    │   ├── Keystore
    │   │   ├── Mutual-Client.crt
    │   │   └── Mutual-Client.key
    │   └── Truststore
    │       └── Mutual-CA.crt
    ├── klient.js
    ├── PaigaldaRakendused.sh
    ├── PaigaldaSerdid.sh
    ├── README.md
    ├── server.js
    └── serverKeys
        ├── Keystore
        │   ├── Mutual-Server.crt
        │   └── Mutual-Server.key
        └── Truststore
            └── Mutual-CA.crt
```
- `CA` - test-CA
- `clientKeys` - kliendi krüptomaterjal 
- `serverKeys` - serveri krüptomaterjal
- `Keystore`  - võtmehoidla (privaatvõtme ja serdi hoidmiseks) 
- `Truststore` - usaldushoidla (CA usaldusahela hoidmiseks) 


