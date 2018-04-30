const http = require('http');
const querystring = require("querystring");
const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
})

async function main() {
    const username = await new Promise((res, rej) => {
        rl.question('username: ', answer => {
            res(answer);
        })
    });

    const password = await new Promise((res, rej) => {
        rl.question('password: ', answer => {
            res(answer);
        });
    });

    rl.close();

    // real password pwd123456
    // SQL injection
    const postData = querystring.stringify({
        username,
        password,
    });

    postLogin(postData);
}

function postLogin (postData) {
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
              console.log(chunk);
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
}

main();
