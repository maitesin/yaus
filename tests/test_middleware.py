from flask import make_response, request, session
from yaus.middleware import verify_url, verify_shortcode, no_recursive_calls_allowed, no_longer_result_url_than_input_url
from werkzeug.exceptions import UnprocessableEntity, NotFound


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
        except NotFound:
            return

        assert False


def test_no_recursive_calls_allowed_valid_url(app):
    with app.test_request_context(data="http://oscarforner.com"):
        session['url'] = "http://oscarforner.com"
        no_recursive_calls_instance = no_recursive_calls_allowed(noop)
        response = no_recursive_calls_instance()

        assert 201 == response.status_code


def test_no_recursive_calls_allowed_invalid_url(app):
    with app.test_request_context(data="https://yaus.dev/wololo"):
        request.url_root = "https://yaus.dev"
        session['url'] = request.url_root + "/wololo"
        no_recursive_calls_instance = no_recursive_calls_allowed(noop)
        try:
            no_recursive_calls_instance()
        except UnprocessableEntity:
            return

        assert False


def test_no_longer_result_url_than_input_url_valid_url(app):
    with app.test_request_context(data="http://oscarforner.com/projects"):
        session['url'] = "http://oscarforner.com/projects"
        no_longer_result_url_than_input_url_instance = no_longer_result_url_than_input_url(noop)
        response = no_longer_result_url_than_input_url_instance()

        assert 201 == response.status_code


def test_no_longer_result_url_than_input_url_invalid_url(app):
    with app.test_request_context(data="https://t.co/w"):
        request.url_root = "https://t.co"
        session['url'] = request.url_root + "/w"
        no_longer_result_url_than_input_url_instance = no_longer_result_url_than_input_url(noop)
        try:
            no_longer_result_url_than_input_url_instance()
        except UnprocessableEntity:
            return

        assert False
