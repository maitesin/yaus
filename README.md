[![Build Status](https://travis-ci.org/maitesin/yaus.svg?branch=master)](https://travis-ci.org/maitesin/yaus)

# YAUS
**YAUS**, **Y**et **A**nother **U**RL **S**hortener

## Features
- [x] Basic URL shortener functionality
- [x] Use a database to store the shortened URLs instead of keeping them in memory
    - [x] Add database error handling
- [x] Check that the data received is a valid URL
- [x] Use a random generated ID for the shortened URLs instead of using an incremental number
    - [x] Split the random generator function from the string trimmer part
    - [x] Add test for the string trimming part
- [x] Check that the shortcode received is up to 8 characters and only contains alphanumeric values
- [x] Find a way to not destroy the database every time the application is deployed
    - [x] Migrate to use PostgreSQL instead of SQLite
- [x] Add a frontend functionality to allow creating the short URLs from the browser itself
- [x] Custom error pages

## TODO
- [ ] Improve the custom error pages to provide more information regarding the reason of the failure
- [ ] Add possibility to block URL to specific domains
    - [ ] Use google safe browsing to obtain a list of flagged domains
- [ ] Make **YAUS** user aware
- [ ] Allow users to create custom named shortened URLs

## How to run YAUS locally
Before running **YAUS** its dependencies have to be met.

### Install dependencies
To install **YAUS**'s dependencies just run the following command:
```bash
pip install -r requirements.txt
```
### Running YAUS
To run **YAUS** run the following command:
```bash
python run.py
```

## How to run YAUS tests
The only requirement is to have tox installed.
```bash
tox
```
