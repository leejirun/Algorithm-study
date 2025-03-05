
const fs = require("fs");

triangular_num = []
for (let i=1; ;i++) {
    let t = (i*(i+1)) /2;
    if (t > 1000) break;
    triangular_num.push(t);
}

function eureka(k) {
    for (let a of triangular_num) {
        for (let b of triangular_num) {
            for (let c of triangular_num) {
                if (a + b+ c=== k) {
                    return 1;
                }
            }
        }
    }
    return 0;
}

const readline = require("readline");
const rl = readline.createInterface({input: process.stdin, output: process.stdout});

let input = [];
rl.on("line", (line) => {
    input.push(line.trim());
}).on("close", () => {
    let T = parseInt(input[0]);
    for (let i = 1; i <= T; i++) {
        let K = parseInt(input[i]); 
        console.log(eureka(K));
    }
    process.exit();
});