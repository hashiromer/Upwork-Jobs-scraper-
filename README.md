# Upwork Jobs Scraper

This project uses **Upwork's internal API** to scrape newly posted jobs. It avoids using a headless browser for two key reasons:

1. **Efficiency**: API requests are far less resource-intensive than browser automation.
2. **Simplicity**: The API returns clean JSON, eliminating the need for HTML parsing and simplifying downstream processing.

---

## Why Golang?

This scraper is written in **Go** instead of Python due to **Upwork's bot detection techniques**, which rely on analyzing TLS signatures of incoming requests ([explained here](https://scrapfly.io/blog/how-to-avoid-web-scraping-blocking-tls/)).

> At the time of development (3 years ago), I could not find an HTTP client in Python that could accurately mimic browser-like TLS signatures. This has since changed as modern Python libraries now support lower-level TLS emulation through bindings to native clients written in Go/C++.

---

## Usage

1. **Add Credentials**
   Create a `.env` file inside the `upwork/` directory. Add your **authorization** and **cookie** headers.

   > 💡 The easiest way to obtain these is by logging into [Upwork.com](https://www.upwork.com), inspecting **network traffic** in DevTools after passing bot checks, and copying the `Authorization` and `Cookie` headers from any authenticated request.

2. **Keep Credentials Fresh**
   These credentials usually expire daily, so you'll need to refresh them if scraping on a regular basis.

3. **Run the Scraper**
   In the `main.go` file, call:

   ```go
   p.Run("<keyword>")
   ```

   * Replace `<keyword>` with the term you want to search.
   * It will fetch the **top 5000 most recent job postings** that match the keyword and save them in a `.jsonl` file.
   * If you pass an empty string (`p.Run("")`), it will fetch the **most recent jobs** regardless of keyword.

> ⚠️ **Important**: Never use your **personal Upwork account** to extract credentials. Doing so **will result in account suspension**.

---

## Contributing

This project is no longer actively maintained, but I occasionally check if it still works (as of **June 27, 2025**, it does).

**Potential contributions:**

* Automating the refresh of `Authorization` and `Cookie` headers.
* Adding support for advanced filters or scraping job details.
