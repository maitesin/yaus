def test_good_redirection(client, entry_url):
    response = client.get(entry_url)

    assert response.status_code == 307
    assert b"Redirecting" in response.data
    assert b"https://oscarforner.com" in response.data


def test_bad_redirection(client):
    response = client.get("/00000001")

    assert response.status_code == 404
    assert b"Page not found" in response.data


def test_invalid_shortcode_too_large(client):
    response = client.get("/123456789")

    assert response.status_code == 404


def test_invalid_shortcode_not_alphanumeric(client):
    response = client.get("/!@#$%^&*")

    assert response.status_code == 422


def test_fail_invalid_url(client):
    response = client.post("/", data="wololo")

    assert response.status_code == 422
    assert "Location" not in response.headers


def test_fail_invalid_url_too_large(client):
    response = client.post("/", data=f"https://{'x' * 10000}.com")

    assert response.status_code == 422
    assert "Location" not in response.headers


def test_fail_invalid_url_too_short(client):
    response = client.post("/", data=f"https://x.com")

    assert response.status_code == 422
    assert "Location" not in response.headers


def test_create_short_url(client):
    response = client.post("/", data="https://wololo.com/somethingelse")

    assert response.status_code == 201
    assert "Location" in response.headers


def test_create_short_url_twice(client):
    response1 = client.post("/", data="https://wololo.com/somethingelse")

    assert response1.status_code == 201
    assert "Location" in response1.headers

    response2 = client.post("/", data="https://wololo.com/somethingelse")

    assert response2.status_code == 201
    assert response1.headers["Location"] == response2.headers["Location"]
