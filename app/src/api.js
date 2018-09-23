const _ = require('lodash');
const host = `http://${window.location.hostname}:8081/`;

export function apiMethod(name) {
    if (name != null && name.length > 0)
        return _.trimEnd(host, '/') + '/' + _.trimStart(name, '/');
    else
        return _.trimEnd(host, '/')
}