# Potential privacy hazard in Fabric chaincodes

This repository tries to demonstrate a potential privacy hazard in Chaincodes due to the execute and then order paradigm of Fabric

## What's going on ?
* in `exploit.go` we write two endpoints for the chaincode `test` and `reset`
* `reset` sets the read token to the initial unused state and sets `a` and `b`
* `test` gets one of either `a` or `b` based on the user's choice
* By design of the function we are allowed to read only `a` or `b` not both.
* Due to the execute first and then order philosophy we are able to read both the secret values if we contact two peers simultaneously
as the commit transaction occurs afterwards.

Screenshot : [Imgur](https://i.imgur.com/YktAb3y.png)

## Potential Hazards
* We believe that a developer who isn't aware of the underlying architecture of Fabric might end up writing code which
  can be exploited.

## Potential Fixes
* The developer of the chaincode should make endpoints(functions) of chaincode lightweight and do read and write in separate
  transactions.

## Steps to Reproduce Demo
* Install the chaincode on two peers
* Instantiate them 
* Call reset on the API to set the read token to "0"
* run the script in this directory

