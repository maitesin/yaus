def test_good_redirection(client, entry_url):
    response = client.get('/00000001')

    assert b'Redirecting' in response.data
    assert b'https://oscarforner.com' in response.data
    assert response.status_code == 307


def test_bad_redirection(client):
    response = client.get('/00000001')

    assert b'404 Not Found' in response.data
    assert response.status_code == 404


def test_fail_invalid_url(client):
    response = client.post('/', data='wololo')

    assert 'Location' not in response.headers
    assert response.status_code == 422


def test_create_short_url(client):
    response = client.post('/', data='https://wololo.com')

    assert 'Location' in response.headers
    assert response.status_code == 201


def test_create_short_url_twice(client):
    response1 = client.post('/', data='https://wololo.com')

    assert 'Location' in response1.headers
    assert response1.status_code == 201

    response2 = client.post('/', data='https://wololo.com')

    assert response1.headers['Location'] == response2.headers['Location']
    assert response2.status_code == 201
