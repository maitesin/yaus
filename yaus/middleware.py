from functools import wraps
from flask import request, abort, session
from validators import url as url_validator


def verify_url(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        url_data = request.form.get("url", request.get_data().decode("utf-8"))
        if not url_validator(url_data) or len(url_data) > 10000:
            abort(422)
        session["url"] = url_data
        return f(*args, **kwargs)

    return decorated_function


def verify_shortcode(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        shortcode = kwargs["shortcode"]
        if not shortcode.isalnum():
            abort(422)
        if len(shortcode) > 8:
            abort(404)
        return f(*args, **kwargs)

    return decorated_function


def no_recursive_calls_allowed(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        url = session["url"]
        if url.startswith(request.url_root):
            abort(422)
        return f(*args, **kwargs)

    return decorated_function


def no_longer_result_url_than_input_url(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        url_length = len(session["url"])
        generated_url_length = len(request.url_root) + 8
        if url_length <= generated_url_length:
            abort(422)
        return f(*args, **kwargs)

    return decorated_function
