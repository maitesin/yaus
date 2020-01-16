from flask import Blueprint, request, make_response, redirect, abort
from short.models import URL
from short.id_generator import id_generator
from short import db
from sqlalchemy import exc
from validators import url as url_validator

short = Blueprint('short', __name__)


@short.route("/", methods=["POST"])
def add_shortcode():
    url = request.get_data()
    if not url_validator(url.decode('utf-8')):
        abort(422)
    key = next(id_generator)
    entry = URL(shortened=key, extended=url)
    db.session.add(entry)
    try:
        db.session.commit()
    except exc.IntegrityError as e:
        assert 'UNIQUE constraint failed' in str(e)
        db.session.rollback()
        entry = URL.query.filter_by(extended=url).first()
        key = entry.shortened
    resp = make_response("Here you have it", 201)
    resp.headers["Location"] = key
    return resp


@short.route("/<string:shortcode>")
def get_url_by_shortcode(shortcode):
    url = URL.query.filter_by(shortened=shortcode).first()
    if url:
        return redirect(url.extended, code=307)
    abort(404)
