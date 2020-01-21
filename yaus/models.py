from yaus import db


class URL(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    shortened = db.Column(db.String, unique=True, nullable=False)
    extended = db.Column(db.String, unique=True, nullable=False)

    def __repr__(self):
        return (
            f"URL(id={self.id}, "
            "shortened={self.shortened}, "
            "extended={self.extended})"
        )
