# Upwork-Jobs-scraper

The code uses Upwork's internal Api to scrape new jobs posted on Upwork. I am not using headless browser due to two reasons.

1. Using headless bowsers is way more resource intensive compared to using an Api.

2. I don't have to deal with HTML parsing, the api returns json which can be directly passed to downstream systems.


# How to use

1. Create a .env file and add the variables.

2. Run go run main.go


# Note

The code uses Golang instead of Python because Upwork filters bots by checking TLS signatures of incoming requests. Unfortunately, I could not
find a way to do it in pure Python because Python is compiled with openssl and popular browsers do not use it. Chrome uses BoringSSl and firefox uses NSS.
These SSL libraries use different extensions and cipher suites which makes detection of TLS level configurations a more robust method to detect bot traffic.

Golang is a more lower level language compared to Python, so it allows changing network level configurations. I am using `cycletls` package in golang which makes spoofing TLS/JA3 fingerprints an easy task.

# How can you contribute?

These are some of the features I think could be useful.
  
- Better error handling with channels  
- Add support for automatic proxy rotation. It can be extremely effective when used in conjunction with go routines. 
- Add Api schema for Upwork Api.
- Add more scrapers, a lot of logic is platform agnostic which could be used to build scrapers for more platforms.
