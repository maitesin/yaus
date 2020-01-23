from yaus import db, create_app
from yaus import Config
import os


class ProdConfig(Config):
    SQLALCHEMY_DATABASE_URI = os.environ.get("DATABASE_URL")


app = create_app(ProdConfig)

db.create_all(app=app)
