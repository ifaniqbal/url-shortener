# URL Shortener

curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://www.google.com"}' http://localhost

ab -n 10000 -c 100 -p payload.json -T 'application/json' http://localhost/