[![Build Status](https://travis-ci.org/maitesin/short.svg?branch=master)](https://travis-ci.org/maitesin/short)

# short
**short** is a URL shortener service

## Features
- [x] Basic URL shortener functionality
- [x] Use a database to store the shortened URLs instead of keeping them in memory
    - [x] Add database error handling
- [x] Check that the data received is a valid URL
- [x] Use a random generated ID for the shortened URLs instead of using an incremental number
    - [ ] Split the random generator function from the string trimmer part
    - [ ] Add test for the string trimming part
- [ ] Check that the shortcode received is up to 8 characters and only contains alphanumeric values
- [ ] Add a frontend functionality to allow creating the short URLs from the browser itself
- [ ] Custom error webpages
- [ ] Make the app user aware
- [ ] Allow users to create custom named shortened URLs

## How to run short locally
Before running **short** its dependencies have to be met.

### Install dependencies
To install **short**'s dependencies just run the following command:
```bash
pip install -r requirements.txt
```
### Running short
To run **short** run the following command:
```bash
python run.py
```

## How to run short tests
The only requirement is to have tox installed.
```bash
tox
```
