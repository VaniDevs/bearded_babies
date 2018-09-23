import { AUTH_LOGIN, AUTH_LOGOUT, AUTH_ERROR, AUTH_CHECK, AUTH_GET_PERMISSIONS } from 'react-admin';
import { apiMethod } from './api'
import decodeJwt from 'jwt-decode';

const roles = {
    1: 'admin',
    2: 'agent'
};

export default (type, params) => {

    if (type === AUTH_LOGIN) {
        const { username, password } = params;
        const request = new Request(apiMethod("login", true), {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        });
        return fetch(request)
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(({ token }) => {
                const decodedToken = decodeJwt(token);
                localStorage.setItem('token', token);
                localStorage.setItem('role', roles[decodedToken.role]);
            });
    }

    if (type === AUTH_LOGOUT) {
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        return Promise.resolve();
    }

    if (type === AUTH_ERROR) {
        const status  = params.status;
        if (status === 401 || status === 403) {
            localStorage.removeItem('token');
            localStorage.removeItem('role');
            return Promise.reject();
        }
        return Promise.resolve();
    }

    if (type === AUTH_CHECK) {
        if (window.location.pathname.startsWith("/appointments")) {
            return Promise.resolve('');
        }
        const token = localStorage.getItem('token');
        if (token) {
            const request = new Request(apiMethod("refresh_token"), {
                method: 'GET',
                headers: new Headers({'Authorization': `Bearer ${token}`})
            });
            return fetch(request)
                .then(response => {
                    if (response.status < 200 || response.status >= 300) {
                        throw new Error(response.statusText);
                    }
                    return response.json();
                })
                .then(({ token }) => {
                    const decodedToken = decodeJwt(token);
                    localStorage.setItem('token', token);
                    localStorage.setItem('role', roles[decodedToken.role]);
                });
        } else {
            return Promise.reject();
        }
    }

    if (type === AUTH_GET_PERMISSIONS) {
        if (window.location.pathname.startsWith("/appointments")) {
            return Promise.resolve('');
        }
        const role = localStorage.getItem('role');
        return role ? Promise.resolve(role) : Promise.reject();
    }

    return Promise.reject('Unknown method');
}