/**
 * Vastastikuse autentimise näiterakendus
 * 
 * Mutual-Client
 */

/* HTTP kliendi teek
  Vt https://www.npmjs.com/package/request 
  http://stackabuse.com/the-node-js-request-module/
*/
const requestModule = require('request');

const fs = require('fs');
const path = require('path');

const certFile = path.resolve(__dirname,
  'clientKeys/Keystore/Mutual-Client.crt');
const keyFile = path.resolve(__dirname,
  'clientKeys/Keystore/Mutual-Client.key');
const caFile = path.resolve(__dirname,
  'clientKeys/Truststore/Mutual-CA.crt');

var options = {
  url: 'https://localhost:5001',
  method: 'GET',
  cert: fs.readFileSync(certFile),
  key: fs.readFileSync(keyFile),
  passphrase: 'changeit',
  ca: fs.readFileSync(caFile)  
}

requestModule(
  options,
  function (error, response, body) {
    if (error) {
      console.log('--- Viga Mutual-Server-i poole pöördumisel: ', error);
      return
    }
    console.log('\n--- Vastus Mutual-Server-st saadud.\n    HTTP staatuskood: clear' +
      response.statusCode);
  }
);