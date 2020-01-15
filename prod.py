from short import db, create_app

app = create_app()

db.create_all(app=app)
