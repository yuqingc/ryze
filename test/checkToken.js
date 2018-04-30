const http = require('http')

const token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUxODk0NjEsImlhdCI6MTUyNTEwMzA2MSwiaXNzIjoiUnl6ZS9nZW5lcmFsIiwidXNyIjoiYWRtaW4ifQ.2DXeZMdrhNBRjFYMhuXFCs0nAVa84jtxdFsNbDYPfpI`;

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