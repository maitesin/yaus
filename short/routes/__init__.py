from flask import Blueprint, request, make_response, redirect, abort
from short.models import URL
from short import id_generator, db

short = Blueprint('short', __name__)


@short.route("/", methods=["POST"])
def add_shortcode():
    url = request.get_data()
    key = next(id_generator)
    entry = URL(shortened=key, extended=url)
    db.session.add(entry)
    db.session.commit()
    resp = make_response("Here you have it", 201)
    resp.headers["Location"] = key
    return resp


@short.route("/<string:shortcode>")
def get_url_by_shortcode(shortcode):
    url = URL.query.filter_by(shortened=shortcode).first()
    if url:
        return redirect(url.extended, code=307)
    abort(404)
