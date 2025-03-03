const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("N을 입력하세요: ", (input) => {
    const N = parseInt(input, 10);
    
    function findGenerator(N) {
        for (let M = 1; M <= N; M++) {
            let sum = M + [...String(M)].reduce((acc, digit) => acc + Number(digit), 0);
    
            if (sum === N) {
                console.log(M);
                rl.close();
                return;
            }
        }
        console.log(0);
        rl.close();
    }

    findGenerator(N);
});
