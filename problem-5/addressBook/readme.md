# addressbook
addressbook is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
addressbook allows the user to perform CRUD operations in the context of contacts via the Command Line Interface.


# CONSENSUS BREAKING CHANGE 

## What does it mean to break consensus
A consensus-breaking change is a modification to the fundamental rules that govern how nodes in a blockchain network validate transactions and reach agreement on the state of the ledger. This type of change can create a division in the network, as nodes may no longer agree on the valid state of the blockchain.

Such a situation often results in a "hard fork," where the network splits into two separate chains that operate under incompatible rules. These separate chains each have its own version of the blockchain history

## Why my change would break consensus 
The change I implemented modifies the structure of the stored contact record by introducing a new field: nickname. This affects both transaction validation and state compatibility, leading to a consensus-breaking scenario.

1. Transaction Validation Changes

The create-contact transaction now requires an additional nickname parameter.
Old nodes do not expect this field and will reject transactions that include it.
This creates inconsistent behavior across the network—new nodes accept the updated transaction format, while old nodes reject it, leading to a potential fork.

2. State Compatibility Issues

Contact entries stored on the blockchain now include the nickname field.
When an old node attempts to process a block containing a contact with this new field, it fails to decode the data because its expected data structure is different.
This results in a state mismatch, where different nodes interpret the blockchain state inconsistently, further splitting the network.

The below commands are updated to accomodate the consensus breaking change.

## Get started

ignite chain build

build command builds the blockchain

ignite chain serve

serve command installs dependencies, builds, initializes, and starts your blockchain in development.

addressbookd keys add [wallet name]

command creates wallet to interact with the blockchain (if needed)



## Creating a Contact 

addressbookd tx addressbook create-contact [first-name] [last-name] [phone] [email] [address] [nickname] [flags]

example usage: 

addressbookd tx addressbook create-contact Hrishikesh Sathyian 87175643 hrishi@gmail.com Yishun BigDawg --from Alice --chain-id addressbook



## Updating a Contact 

addressbookd tx addressbook update-contact [first-name] [last-name] [phone] [email] [address] [nickname] [0] [flags]

example usage: 

addressbookd tx addressbook create-contact Hrishikesh Sathyian 87175643 hrishi@gmail.com Yishun SmallDawg 0 --from Alice --chain-id addressbook


## Deleting a Contact 

addressbookd tx addressbook delete-contact [id]

example usage: 

addressbookd tx addressbook delete-contact 0 --from Alice --chain-id addressbook



## Querying Contacts 

addressbookd q addressbook list-contact 

list-contact displays all contacts in the blockchain

 addressbookd q addressbook show-contact [id]

show-contact queries the contact with the specified id number




## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)