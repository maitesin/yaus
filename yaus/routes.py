from flask import Blueprint, session, make_response, redirect, abort, render_template, flash, request, Markup
from yaus.models import URL
from yaus.id_generator import id_generator
from yaus.middleware import verify_url, verify_shortcode
from yaus import db
from sqlalchemy import exc

yaus = Blueprint('yaus', __name__)


@yaus.route("/", methods=["GET"])
def home():
    return render_template('home.html')


@yaus.route("/", methods=["POST"])
@verify_url
def add_shortcode():
    url = session['url']
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
    flash(Markup(f'Short URL created <a href="{request.url_root + key}">{request.url_root + key}</a>'), 'success')
    resp = make_response(render_template('home.html'), 201)
    resp.headers["Location"] = key
    return resp


@yaus.route("/<string:shortcode>")
@verify_shortcode
def get_url_by_shortcode(shortcode):
    url = URL.query.filter_by(shortened=shortcode).first()
    if url:
        return redirect(url.extended, code=307)
    abort(404)
