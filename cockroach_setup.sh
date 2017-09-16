cockroach start --insecure --background
cockroach sql --insecure -e 'CREATE DATABASE wubalubadubdub;'
cockroach sql --insecure -e 'CREATE TABLE wubalubadubdub.words (id SERIAL PRIMARY KEY, phrase STRING, file STRING, startPhrase TIMESTAMP, endPhrase TIMESTAMP, nextPhrase TIMESTAMP);'
cockroach sql --insecure -e 'SELECT * FROM wubalubadubdub.words;'
