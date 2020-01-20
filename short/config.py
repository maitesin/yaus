import secrets


class Config:
    SQLALCHEMY_DATABASE_URI = "sqlite:///short.db"
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    SECRET_KEY = secrets.token_hex(16)
