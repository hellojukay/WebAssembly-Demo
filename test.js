"use strict";
const fs = require('fs');
const { execSync } = require('child_process');
const wasmExec = execSync('go env GOROOT').toString().replace(/\n/g, '') + '/misc/wasm/wasm_exec.js';
require(wasmExec);

const buf = fs.readFileSync('main.wasm')
const go = new g.Go();

WebAssembly.Instance(buf,go.importObject).then(result => {
    go.run(result.Instance)
    console.log(addFun)
})