from flask_wtf import FlaskForm
from wtforms import StringField
from wtforms.validators import URL, DataRequired, Length


class URLShortenerForm(FlaskForm):
    url = StringField(
        "url",
        render_kw={
            "placeholder": "URL to make shorter",
            "class": "form-control col-md-10",
        },
        validators=[URL(), DataRequired(), Length(min=1, max=10000)],
    )
