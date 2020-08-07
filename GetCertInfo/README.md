GoTools on sertide kasutuse kontrollija.

GoTools kontrollib, kas rakenduste serdid on õigesti seatud, tehes päringud rakenduste HTTPS otspunktide vastu ja kuvadev iga leitud serdi väljaandja (Issuer), subjekti (Subject) ja Subject Alternate Names tähtsamad elemendid.

Nt:

```
$ go run . -conf ../GoTools-config/config.json
loadConf:
--- Loen seadistuse failist: ../GoTools-config/config.json
    Seadistus laetud:
    main.Config{ClientCert:"../GoTools-config/https.crt", ClientKey:"../GoTools-config/https.key" <..>
loadVObjects:
--- Loen kontrolli-URL-id
    Loetud kontrolliobjektid:
    Moodul <..> õlg 1: https://<hostinimi>/health
    Moodul <..> õlg 2: https://<hostinimi>/health
    Moodul <..> õlg 1: https://<hostinimi>/health
    <..>
--- Kontrollin masinat:
    Moodul <..> õlg 1
    https://<hostinimi>/health
    Issuer:
    O=[<..>], CN=<hostinimi>
    Subject:
    O=[<..>], CN=<hostinimi>
    Subject Alternate Names:
    [<hostinimi> <hostinimi> <hostinimi>]
<..>
```

Paigaldamine: 1) klooni repo; 2) moodusta seadistuse kaust, nt `GoTools-config`; 3) Moodusta kausta seadistusfail (näidis - `config.json`) ja fail kontrollitavate masinate URL-de loeteluga (näidis - `VObjects.json`); 4) Genereeri võtmed; 5) Sea väärtused seadistusfailis.

Käivitamine:

```
go run . -conf <seadistusfail>
```

