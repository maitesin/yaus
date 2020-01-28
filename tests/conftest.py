import os
import tempfile
import pytest

from yaus import create_app, db


@pytest.fixture
def app():
    app = create_app()
    db_fd, app.config["DATABASE"] = tempfile.mkstemp()
    db_uri = f"sqlite:///{app.config['DATABASE']}"
    app.config["SQLALCHEMY_DATABASE_URI"] = db_uri
    app.config["TESTING"] = True

    db.create_all(app=app)

    yield app

    os.close(db_fd)
    os.unlink(app.config["DATABASE"])


@pytest.fixture
def client(app):
    with app.test_client() as client:
        yield client


@pytest.fixture
def entry_url(client):
    return client.post("/", data="https://oscarforner.com/projects").headers["Location"]
