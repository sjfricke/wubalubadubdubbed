# Wubalubadubdubbed

This is a project to take 3 season of Rick and Morty and index all the words spoken in a database

# UPDATE

Due to the 500+ visitors in one month driving my electricty cost, I sadly have taken down the live demo. I do plan to update repo to allow anyone who has server space to host it instead.

# How the wubalub gets a dubdubbed

- First all episodes are ingested through IBM Watson to get the speech to text
- All data is stored in the CockroachDB database
- Using Go on a post request of searched words it will attempt to find the words and contruct a dubbed montage of Rick and Morty clips to that phrase

# Bugs

- A LOT
-- Todo: Add them all :(
