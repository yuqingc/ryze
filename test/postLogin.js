const http = require('http');
const querystring = require("querystring");

const postData = querystring.stringify({
    username: 'Matt'
});

const req = http.request(
    {
        port: '8080',
        method: 'POST',
        path: '/api/login',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Content-Length': Buffer.byteLength(postData)
        }
    },
    (res) => {
        console.log(`STATUS: ${res.statusCode}`);
        console.log(`HEADERS: ${JSON.stringify(res.headers)}`);
        res.setEncoding('utf8');
        res.on('data', (chunk) => {
          console.log(`BODY: ${chunk}`);
        });
        res.on('end', () => {
          console.log('No more data in response.');
        });
      }
);

req.on('error',function(err){  
    console.error(err);  
});

req.write(postData);

req.end();