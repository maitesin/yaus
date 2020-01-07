def test_good_redirection(client, entry_url):
    response = client.get('/00000001')

    assert b'Redirecting' in response.data
    assert b'https://oscarforner.com' in response.data
    assert response.status_code == 307


def test_bad_redirection(client):
    response = client.get('/00000001')

    assert b'404 Not Found' in response.data
    assert response.status_code == 404


def test_create_short_url(client):
    response = client.post('/', data='https://wololo.com')

    assert 'Location' in response.headers
    assert response.status_code == 201
