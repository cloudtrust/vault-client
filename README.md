# vault-client

Vault client is a client that can:

- **read**
- **write**
- **create a new policy**
- **create a token** (associated to a policy) in Vault. 


## Launch 

```bash
    import vaultClient "github.com/cloudtrust/vault-client/client"
	var vClient vaultClient.Client
	{
		var err error
		vClient, err = vaultClient.NewClient(vaultToken, vaultURL)
		if err != nil {
			panic(err)
		}
	}

``` 

## Configuration 

Vault needs to be set up. The vaultToken used should allow the vault client to perform all his operations.  

## Usage - examples

The methods of this client follow the syntax of the client provided by Vault. 

*read*

```bash
secret, errRead = vClient.Read(pathKey, token)
```
The vault client reads the information stored on ```pathKey``` in Vault.


*write*

```bash
_, errWrite = vClient.Write(pathKey, map[string]interface{}{"key": keyValue}, token)
```
The vault client writes on the path ```pathKey``` the key/value ```"key": keyValue```.

*create policy*

```bash
err = vClient.CreatePolicy(pathPolicy, "writekey", policyName)
```
In order to create a policy, the vault client needs to specify the path of the policy, the name of the policy and the role of that policy. In this example, a policy that gives the right to write a key in Vault is created.

The existing roles are *writekey*, *readkey*, *createkey*, *exportkey*, *encrypt* and *decrypt*. These correspond to the functionality needed by github.com/cloudtrust/vault-bridge .  

By default, the ttl of the policy is of 1 hour. 

*create token*

```bash
token, errToken = vClient.CreateToken(policyName)
```
The vault client creates a Vault token associated to the policy with the name ```policyName```.


