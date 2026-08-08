import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';

const fallbackRate = new Rate('fallback_triggered');

export const options = {
    vus: 1,
    iterations: 20,
};

const BASE_URL = 'http://localhost:8080/api/call?slow=true';

export default function () {
    let res = http.get(BASE_URL);

    check(res, {
        'is status 200': (r) => r.status === 200,
    });

    const isFallback = res.body.includes('Fallback');
    fallbackRate.add(isFallback);

    // Виводимо статус та повністю тіло відповіді (body), яке прийшло від сервера
    console.log(`[Ітерація ${__ITER + 1}] Статус: ${res.status} | Відповідь: ${res.body}`);

    sleep(0.5);
}