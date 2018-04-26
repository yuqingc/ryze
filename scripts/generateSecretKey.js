const crypto = require('crypto');
const fs = require('fs');
const path = require('path')

const secretKey = crypto.randomBytes(256).toString('hex') + '\n';
fs.writeFile(path.resolve('/home/matt/secret_key'), secretKey, function(err) {
  if (err) throw err;
  console.log('Written!');
});
