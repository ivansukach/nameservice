# Initialize configuration files and genesis file
# moniker is the name of your node
nsd init <moniker> --chain-id namechain

# Configure your CLI to eliminate need to declare them as flags
nscli config chain-id namechain
nscli config output json
nscli config indent true
nscli config trust-node true

# We'll use the "test" keyring backend which save keys unencrypted in the configuration directory of your project (defaults to ~/.nsd). You should **never** use the "test" keyring backend in production. For more information about other options for keyring-backend take a look at https://docs.cosmos.network/master/interfaces/keyring.html
nscli config keyring-backend test

# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
nscli keys add jack

# Copy the `Address` output here and save it for later use
nscli keys add alice

# Add both accounts, with coins to the genesis file
nsd add-genesis-account $(nscli keys show jack -a) 1000nametoken,100000000stake
nsd add-genesis-account $(nscli keys show alice -a) 1000nametoken,100000000stake

# The "nscli config" command saves configuration for the "nscli" command but not for "nsd" so we have to
# declare the keyring-backend with a flag here
nsd gentx --name jack <or your key_name> --keyring-backend test
