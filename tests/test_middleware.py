from flask import make_response
from yaus.middleware import verify_url, verify_shortcode
from werkzeug.exceptions import UnprocessableEntity


def noop(*args, **kwargs):
    return make_response("Here you have it", 201)


def test_verify_url_successful(app):
    with app.test_request_context(data="http://oscarforner.com"):
        verify_url_instance = verify_url(noop)
        response = verify_url_instance()

        assert 201 == response.status_code


def test_verify_url_invalid_url(app):
    with app.test_request_context(data="wololo"):
        verify_url_instance = verify_url(noop)
        try:
            verify_url_instance()
        except UnprocessableEntity:
            return

        assert False


def test_verify_shortcode_successful(app):
    with app.test_request_context():
        verify_shortcode_instance = verify_shortcode(noop)
        response = verify_shortcode_instance(shortcode="abcd1234")

        assert 201 == response.status_code


def test_verify_shortcode_invalid_shortcode_not_alphanumeric(app):
    with app.test_request_context():
        verify_shortcode_instance = verify_shortcode(noop)
        try:
            verify_shortcode_instance(shortcode="!@#$%^&*")
        except UnprocessableEntity:
            return

        assert False


def test_verify_shortcode_invalid_shortcode_too_length(app):
    with app.test_request_context():
        verify_shortcode_instance = verify_shortcode(noop)
        try:
            verify_shortcode_instance(shortcode="123456789")
        except UnprocessableEntity:
            return

        assert False
