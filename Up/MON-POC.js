'use strict';

var perioodid = [
  "10s", "1m", "5m", "15m", "30m", "1h", "2h", "6h", "12h", "1d", "2d",
  "1w", "2w", "4w", "8w", "1y", "2y"
];

var pi = 5; // Indeks massiivis perioodid.

var 
  algusEpoch, // Unix ajamärgendina, arvuna
  loppEpoch, // Unix ajamärgendina, arvuna
  sammEpoch; // sekunditena, arvuna.

var urliAlgus;

function alusta() {
  console.log('Alustan...');
  $('#Pikkus').text('1h');
  seaNupukasitlejad();
  urliAlgus = 'http://192.168.56.1:9090/api/v1/query_range' +
    '?query=up';
  pariJaKuva();
}

function pariJaKuva() {
  // Eemalda vana andmed
  $("#Tulemus").empty();
  // Valmista ajaparameetrid ette.
  var pp = perioodid[pi];
  var vaartus = parseInt(pp.slice(0, pp.length - 1));
  var yhik = pp.slice(-1);
  var lopp = moment(); // Jooksev aeg.
  var algus = moment().subtract(vaartus, yhik);
  // Arvuta samm, jagades perioodi 20-ks osaks.
  // Samm sekundites
  sammEpoch = moment.duration(vaartus, yhik).as('seconds') / 20;
  // Algus, lõpp ja samm stringina
  var alguss = algus.format('X');
  var lopps = lopp.format('X');
  var samms = sammEpoch.toString() + 's';
  algusEpoch = parseInt(alguss);
  loppEpoch = parseInt(lopps);
  // Päri andmed Prometheus-lt.
  var url = urliAlgus +
    '&start=' + alguss +
    '&end=' + lopps +
    '&step=' + samms;
  fetch(url)
    .then(res => res.json())
    .then((data) => {
      // $('#Tulemus').text(JSON.stringify(data, null, 4));
      if (data.status == "success") {
        kuvaAndmed(data);
      } else {
        $('#Teade').text('Andmepäring ebaõnnestus');
      }
    }).catch(err => {
      $('#Teade').text('Andmepäring ebaõnnestus');
      console.error(err);
    });

}

function kuvaAndmed(data) {
  var m = data.data.result; // Andmemaatriks

  // Täida lüngad aegridades nullidega.
  for (var i = 0; i < m.length; i++) {
    var nullidega = [];
    var pos = 0;
    for (var t = algusEpoch; t <= loppEpoch; t += sammEpoch) {
      // Allow for floating point inaccuracy.
      if (pos < m[i].values.length && m[i].values[pos][0] < t + sammEpoch / 100) {
        nullidega.push(m[i].values[pos]);
        pos++;
      } else {
        nullidega.push([t, '0']);
      }
    }
    m[i].values = nullidega;
  }

  m.forEach(element => {
    // Lisa uus rida.
    var div = $("<div></div>");
    $("<span></span>").addClass('otspunkt').text(element.metric.instance).appendTo(div);
    $("#Tulemus").append(div);
    var v = element.values; // Väärtuste vektor
    v.forEach(e2 => {
      $("<span></span>").text(e2[1] + ' ').appendTo(div);
    });
  });
}

function seaNupukasitlejad() {
  $('#Pikemaks').click(() => {
    if (pi + 1 < perioodid.length) {
      pi++;
      $('#Pikkus').text(perioodid[pi]);
      pariJaKuva();
    }
  });
  $('#Lyhemaks').click(() => {
    if (pi > 0) {
      pi--;
      $('#Pikkus').text(perioodid[pi]);
      pariJaKuva();
    }
  });
}

// Märkmed
// Reading JSON with Fetch API
// http://zetcode.com/javascript/jsonurl/