# addressbook
**addressbook** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
**addressbook** allows the user to perform CRUD operations in the context of contacts via the Command Line Interface.

## Get started
```
ignite chain build
```
`build` command builds the blockchain
```
ignite chain serve
```
`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.
```
addressbookd keys add [wallet name]
```
command creates wallet to interact with the blockchain (if needed)



## Creating a Contact 
```
addressbookd tx addressbook create-contact [first-name] [last-name] [phone] [email] [address] [flags]
```
example usage: 
```
addressbookd tx addressbook create-contact Hrishikesh Sathyian 87175643 hrishi@gmail.com Yishun --from Alice --chain-id addressbook
```

## Creating a Contact 
```
addressbookd tx addressbook create-contact [first-name] [last-name] [phone] [email] [address] [flags]
```
example usage: 
```
addressbookd tx addressbook create-contact Hrishikesh Sathyian 87175643 hrishi@gmail.com Yishun --from Alice --chain-id addressbook
```

## Updating a Contact 
```
addressbookd tx addressbook update-contact [first-name] [last-name] [phone] [email] [address] [id] [flags]
```
example usage: 
```
addressbookd tx addressbook create-contact Hrishikesh Sathyian 87175643 hrishi@gmail.com Yishun 0 --from Alice --chain-id addressbook
```

## Deleting a Contact 
```
addressbookd tx addressbook delete-contact [id]
```
example usage: 
```
addressbookd tx addressbook delete-contact 0 --from Alice --chain-id addressbook
```


## Querying Contacts 
```
addressbookd q addressbook list-contact 
```
`list-contact` displays all contacts in the blockchain
```
 addressbookd q addressbook show-contact [id]
```
`show-contact` queries the contact with the specified id number




## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
