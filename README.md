# go-vault-demo

## Setup Environment

### Setup Vault Server
```sh
$ docker compose up -d
```

```
[+] Running 1/1
 ⠿ Container vault-server  Started 
```

##### Check Vault Server Status
```sh
$ docker exec -ti vault-server vault status
```

```
Key                Value
---                -----
Seal Type          shamir
Initialized        false
Sealed             true
Total Shares       0
Threshold          0
Unseal Progress    0/0
Unseal Nonce       n/a
Version            1.17.2
Build Date         2024-07-05T15:19:12Z
Storage Type       file
HA Enabled         false
```

### Initialize Vault Server
```sh
$ docker exec -ti vault-server vault operator init
```

```
Unseal Key 1: uIwFsSWaIehNQNyhUAyyid+hHt/ARC3Qh9s2nAFpjEFS
Unseal Key 2: lcJge7im9d04URIv9desLqnBu9wjviQfmvfQQzUopHqP
Unseal Key 3: rM2X7dkg0LSRw+9T4s2yIXz4sCs+EganF5hBhs/umHxY
Unseal Key 4: xyDZkpL441cWftzrOgsEFm2YXSkq9IBFhhUwNLOFM9/w
Unseal Key 5: MB0sHHqGPlfRLIKWnFclkFErAc/AUrQcjV8/+bfbRm2U

Initial Root Token: hvs.y1oxdymlUOxptFEsVPUnmJWi
```

### Unseal Vault using the Unseal Key

```sh
$ docker exec -ti vault-server vault operator unseal <unseal-key>
```
> Repeat this process 3 times using a different unseal key each time.

##### Check Vault Server Status
```sh
$ docker exec -ti vault-server vault status
```

```
Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    5
Threshold       3
Version         1.17.2
Build Date      2024-07-05T15:19:12Z
Storage Type    file
Cluster Name    vault-cluster-63434491
Cluster ID      55d45299-71cf-52f6-dde6-73e451002b39
HA Enabled      fals
```

### Setup Environment Variables
```sh
$ export VAULT_ADDR=http://127.0.0.1:8200
$ export VAULT_TOKEN=<initial_root_token>
```

## Setup Vault Secret Engine

### Create K/V secrets engine

```sh
$ vault secrets enable -path=secret -version=2 kv
```


### Create transit secrets engine

```sh
$ vault secrets enable -path=transit transit
```

##### Check Vault Secret Engines

```sh
$ vault secrets list
```
```
Path          Type         Accessor              Description
----          ----         --------              -----------
cubbyhole/    cubbyhole    cubbyhole_2e25bc33    per-token private secret storage
identity/     identity     identity_9ee1c816     identity store
secret/       kv           kv_b713ca64           n/a
sys/          system       system_66642b49       system endpoints used for control, policy and debugging
transit/      transit      transit_6b91e1fa      n/a
```

## Install Golang Module 

```sh
$ go mod tidy
```

## Execute the Main Program

```sh
$ go run main.go
```

```
2024/08/31 15:33:30 Super secret password written successfully to the vault.
Encrypted data: vault:v1:QouEKTS6KNObAOo5Q/hkhhuzWdgCJ0I4BPfcDVS1yeZqP24ZDsP4MbuD
Decrypted data: my secret data
```
