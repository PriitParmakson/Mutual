/**
 * Vastastikuse autentimise näiterakendus
 * 
 * Mutual-Server
 */

var express = require('express');
var fs = require('fs');
var https = require('https');

var app = express();

var options = {
  // Serveri privaatvõti
  key: fs.readFileSync('serverKeys/Keystore/Mutual-Server.key'),
  // Serveri sert
  cert: fs.readFileSync('serverKeys/Keystore/Mutual-Server.crt'),
  // Serveri usaldushoidla (trust store) - CA-fail
  ca: fs.readFileSync('serverKeys/Truststore/Mutual-CA.crt')
};

/**
 * Päringukäsitleja, mis logib kliendi saadetud serdi
 */
app.use(function (req, res, next) {
  var log = new Date() + ' ' + req.connection.remoteAddress + ' ' + req.method + ' ' + req.url;
  var cert = req.socket.getPeerCertificate();
  if (cert.subject) {
    log += ' ' + cert.subject.CN;
  }
  console.log(log);
  next();
});

/**
 * Päringukäsitleja, mis kontrollib kliendi saadetud serti
 */
app.use(function (req, res, next) {
  if (!req.client.authorized) {
    return res.status(401).send('User is not authorized');
  }
  next();
});

/**
 * Päringukäsitleja, mis koostab ja saadab vastuse
 */
app.use(function (req, res, next) {
  res.writeHead(200);
  res.end("Klient autenditud\n");
  next();
});

options.requestCert = true;

var listener = https.createServer(options, app)
  .listen(5001, function () {
    console.log('Mutual-Server kuulab pordil: ' +
      listener.address().port);
  });