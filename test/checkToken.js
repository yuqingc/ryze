const http = require('http')

const token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjQ4NDE1MDIsImlhdCI6MTUyNDc1NTEwMiwiaXNzIjoiUnl6ZS9nZW5lcmFsIiwidXNyIjoiYWRtaW4ifQ.OLyNX2mVsxZlY50y5RHb1eWx3vsomYfr1HjAMl-zmdQ`;

const badtoken = `fdafadadfasdf`;

const req = http.request(
    {
        port: '8080',
        method: 'GET',
        path: '/api/varify_token',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': `Bearer  ${token}`
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
)

req.on('error',function(err){  
    console.error(err);  
});

req.end();