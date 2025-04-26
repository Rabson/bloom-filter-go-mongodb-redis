
import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://localhost:8080/check-username';
  const payload = JSON.stringify({ username: 'testuser' });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
}
