cockroach sql --insecure -e 'SELECT COUNT(*) FROM wubalubadubdub.words;'

cockroach sql --insecure -e 'SELECT COUNT (DISTINCT phrase) FROM wubalubadubdub.words;'
