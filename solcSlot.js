var solc = require('solc');
solidityCode = `// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.13 and less than 0.9.0
pragma solidity ^0.8.13;

contract HelloWorld {
    string public greet;

    constructor(string memory g) {
        greet = g;
    }
}
`
var input = {
  language: 'Solidity',
  sources: {
    'test.sol': {
      content: solidityCode//'pragma solidity ^0.8.20; contract C { function f() public { } }'
    }
  },
  settings: {
    outputSelection: {
      '*': {
        '*': ['*']
      }
    }
  }
};

var output = JSON.parse(solc.compile(JSON.stringify(input)));
//const interface = output.contracts['inbox.sol'].Inbox.abi;
// `output` here contains the JSON output as specified in the documentation

for (var contractName in output.contracts['test.sol']) {
  console.log(
   
      JSON.stringify(output.contracts['test.sol'][contractName].storageLayout)
  );
  
}

