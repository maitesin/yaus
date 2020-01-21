from flask import Blueprint, request, make_response, redirect, abort, render_template
from short.models import URL
from short.id_generator import id_generator
from short.middleware import verify_url, verify_shortcode
from short import db
from sqlalchemy import exc

short = Blueprint('short', __name__)


@short.route("/", methods=["GET"])
def home():
    return render_template('home.html')


@short.route("/", methods=["POST"])
@verify_url
def add_shortcode():
    url = request.get_data()
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
    resp = make_response(render_template('home.html'), 201)
    resp.headers["Location"] = key
    return resp


@short.route("/<string:shortcode>")
@verify_shortcode
def get_url_by_shortcode(shortcode):
    url = URL.query.filter_by(shortened=shortcode).first()
    if url:
        return redirect(url.extended, code=307)
    abort(404)
