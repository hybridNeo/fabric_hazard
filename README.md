# Potential privacy hazard in Fabric chaincode
This repository tries to demonstrate a potential privacy hazard in Chaincodes due to the execute and then order paradigm of Fabric

## What's going on ?
* This is an example chaincode that someone could reasonably try to write. It exposes one endpoint, called "ReadAorB". 
This allows the caller to read one of the values in the key value store. However, the function has a side effect, which is to mark a special flag changing "readToken = True" 
the first time it is called. The effect of this flag is that the caller can read one of the values, A or B, but once they've picked, they cannot read the other one. 
The overall effect of this chaincode is that it lets the client read one value but not both, enforcing a privacy policy.
The goal of this post is to illustrate how this straightforward contract actually falls for a programming hazard when writing chaincodes, which would let a client violate 
the intended security policy.
* By design of the function we are allowed to read only `a` or `b` not both.
* Due to the execute first and then order philosophy we are able to read both the secret values if we contact two peers simultaneously
    as the commit transaction occurs afterwards.

    Screenshot : [Imgur](https://i.imgur.com/YktAb3y.png)

## Potential Hazards
This illustration of reading A or B is just a toy example, but enforcing read-access control in this programmatic way is representative of what chaincode authors may write in 
many scenarios:
* pay for access: a client may purchase a "token" to read one record from the database, e.g. to download an ebook. speculative reads would let a client effectively double 
  spend that token
* access logging: a pragmatic approach to privacy is to allow clients fairly broad access to read values (e.g., doctors can access any medical record in a hospital),
  but to store an audit trail log every time a client reads a new value. That way if a client is reading data they shouldn't, at least they can be tracked 
  down. https://www.hipaaformsps.com/hipaa-access-logs-audits/  
  However, this speculative read hazard could allow a client to read a value but bypass the audit log, again violating the intended policy
* zero knowledge proofs, oblivious transfer, or related cryptography gadgets: 
 this one is a bit more technical, but a pattern that often arises in cryptography is that a client (a "proof verifier") must be allowed to choose one value to read, 
 but not more than one. For a fairly high level explanation, see here: https://blog.cryptographyengineering.com/2014/11/27/zero-knowledge-proofs-illustrated-primer/

## Potential Fixes
* The developer of the chaincode should make endpoints(functions) of chaincode lightweight and do read and write in separate
    transactions.

## Steps to Reproduce Demo
* Install the chaincode on two peers
* Instantiate them 
* run the script in this directory

