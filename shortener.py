from flask import Flask, request, make_response, redirect, abort
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///short.db'
db = SQLAlchemy(app)


class URL(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    shortened = db.Column(db.String, unique=True, nullable=False)
    extended = db.Column(db.String, unique=True, nullable=False)

    def __repr__(self):
        return f'URL(id={self.id}, shortened={self.shortened}, extended={self.extended})'


def get_id_generator():
    id = 1
    while True:
        yield f'{id}'.zfill(8)
        id += 1


id_generator = get_id_generator()


@app.route('/', methods=['POST'])
def add_shortcode():
    url = request.get_data()
    key = next(id_generator)
    entry = URL(shortened=key, extended=url)
    db.session.add(entry)
    db.session.commit()
    resp = make_response('Here you have it', 201)
    resp.headers['Location'] = key
    return resp


@app.route('/<string:shortcode>')
def get_url_by_shortcode(shortcode):
    url = URL.query.filter_by(shortened=shortcode).first()
    if url:
        return redirect(url.extended, code=307)
    abort(404)
