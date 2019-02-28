#!/bin/bash

CMD_STMT="peer chaincode invoke -C mychannel -o orderer.example.com:7050 -n exploit -c '{\"Args\":[\"test\","
END_STMT="]}'"
CONFIG_PATH=conf/hbavss.local.ini



# Generate config file locally
set -x
tmux new-session    "CORE_PEER_ADDRESS=peer0.org1.example.com:7051 ${CMD_STMT}  "\"a\""  ${END_STMT}; sh" \; \
     splitw -v -p 50 "CORE_PEER_ADDRESS=peer1.org1.example.com:7051 ${CMD_STMT}  "\"b\"" ${END_STMT}; sh" ; 
