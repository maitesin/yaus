from flask import Flask, request, make_response, redirect, abort

app = Flask(__name__)
m = {}


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
    global m
    m[key] = url
    resp = make_response('Here you have it', 201)
    resp.headers['Location'] = key
    return resp


@app.route('/<string:shortcode>')
def get_url_by_shortcode(shortcode):
    if shortcode in m:
        return redirect(m[shortcode], code=307)
    abort(404)
