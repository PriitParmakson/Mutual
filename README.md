# Mutual

Moodusta masinas kaustastruktuur:

```
├── CA
│   ├── TARA-Mutual-CA.crt
│   ├── TARA-Mutual-CA.key
│   ├── TARA-Mutual-CA.srl
│   ├── TARA-Mutual-Client.crt
│   ├── TARA-Mutual-Client.csr
│   ├── TARA-Mutual-Client.key
│   ├── TARA-Mutual-Server.crt
│   ├── TARA-Mutual-Server.csr
│   └── TARA-Mutual-Server.key
├── clientKeys
│   ├── Keystore
│   │   ├── TARA-Mutual-Client.crt
│   │   └── TARA-Mutual-Client.key
│   └── Truststore
│       └── TARA-Mutual-CA.crt
├── serverKeys
│   ├── Keystore
│   │   ├── TARA-Mutual-Server.crt
│   │   └── TARA-Mutual-Server.key
│   └── Truststore
│       └── TARA-Mutual-CA.crt
└── TARA-Mutual
    ├── GenereeriVotmed.sh
    ├── klient.js
    ├── README.md
    └── server.js
```

 kaust | otstarve
-------|----------
`CA`   | test-CA
`clientKeys` | kliendi krüptomaterjal 
`serverKeys`  | serveri krüptomaterjal
`Keystore`  | võtmehoidla (privaatvõtme ja serdi hoidmiseks) 
`Truststore` | usaldushoidla (CA usaldusahela hoidmiseks) 
`TARA-Mutual` | kliendi ja serveri kood (repo `https://github.com/PriitParmakson/TARA-Mutual` peegeldus)


