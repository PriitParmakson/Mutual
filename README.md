# Mutual

Moodusta masinas kaustastruktuur:

├── CA
├── clientKeys
│   ├── Keystore
│   └── Truststore
├── serverKeys
│   ├── Keystore
│   └── Truststore
└── TARA-Mutual
    ├── klient.js
    ├── README.md
    └── server.js

 kaust | otstarve
-------|----------
`CA`   | test-CA
`clientKeys` | kliendi krüptomaterjal 
`serverKeys`  | serveri krüptomaterjal
`Keystore`  | võtmehoidla (privaatvõtme ja serdi hoidmiseks) 
`Truststore` | usaldushoidla (CA usaldusahela hoidmiseks) 
`TARA-Mutual` | kliendi ja serveri kood (repo `https://github.com/PriitParmakson/TARA-Mutual` peegeldus)


