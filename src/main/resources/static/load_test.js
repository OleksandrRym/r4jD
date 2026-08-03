import http from 'k6/http';

export const options = {
    vus: 300000,
    duration: '30s',
};

export default function () {
    const res = http.get('http://localhost:8080/api/pay');
    console.log(`Status: ${res.status}`);
}