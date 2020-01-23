from functools import wraps
from flask import request, abort, session
from validators import url as url_validator


def verify_url(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        url_data = request.form.get("url", request.get_data().decode("utf-8"))
        if not url_validator(url_data):
            abort(422)
        session["url"] = url_data
        return f(*args, **kwargs)

    return decorated_function


def verify_shortcode(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        shortcode = kwargs["shortcode"]
        if len(shortcode) > 8 or not shortcode.isalnum():
            abort(422)
        return f(*args, **kwargs)

    return decorated_function
