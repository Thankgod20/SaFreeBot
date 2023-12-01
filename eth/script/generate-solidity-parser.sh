#!/bin/bash

set -e

# This setup assumes you have antlr4 installed on your system
# as described at https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
java \
    -Xmx500M \
    -cp "/usr/local/lib/antlr-4.12.0-complete.jar:$CLASSPATH" \
    org.antlr.v4.Tool \
    -Dlanguage=Go solidity-antlr4/Solidity.g4 \
    -package solidityparser \
    -o parser
mv parser/solidity-antlr4/* solidityparser

# Cleanup
rm -rf parser