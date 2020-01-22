from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from yaus.config import Config

db = SQLAlchemy()


def create_app(config_class=Config):
    app = Flask(__name__)
    app.config.from_object(config_class)

    db.init_app(app)

    from yaus.routes import yaus as yaus_routes

    app.register_blueprint(yaus_routes)

    return app
