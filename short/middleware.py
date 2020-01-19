from functools import wraps
from flask import request, abort
from validators import url as url_validator


def verify_url(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        url_data = request.get_data()
        if not url_validator(url_data.decode('utf-8')):
            abort(422)
        return f(*args, **kwargs)
    return decorated_function


def verify_shortcode(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        shortcode = kwargs['shortcode']
        if len(shortcode) > 8 or not shortcode.isalnum():
            abort(422)
        return f(*args, **kwargs)
    return decorated_function
