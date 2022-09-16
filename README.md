# Upwork-Jobs-scraper

The code uses Upwork's internal Api to scrape new jobs posted on Upwork. I am not using headless browser due to two reasons.

1. Using headless bowsers is way more resource intensive compared to using an Api.

2. I don't have to deal with HTML parsing, the api returns json which can be directly passed to downstream systems.

# Note

I had to write the script in Golang instead of Python because Upwork filters bots by checking TLS signatures of incoming requests. Unfortunately, I could not
find a way to do it in pure Python because Python is compiled with openssl and popular browsers do not use it. Chrome uses BoringSSl and firefox uses NSS.
These SSL libraries use different extensions and cipher suites which makes detection of TLS level configurations a more robust method to detect bot traffic.

Golang is a more lower level language compared to Python, so it allows changing network level configurations. I am using `cycletls` package in golang which makes spoofing TLS/JA3 fingerprints an easy task.

# How can you contribute?

The code as of now is written as a throwaway script because I wrote it to scrape some data once. If you intend to include it in your workflow or integrate it in larger codebase, you can contribute back by refactoring the code since I don't know golang at all, I had to Google fu everything.
