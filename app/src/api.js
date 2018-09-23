const _ = require('lodash');
const host = `http://${window.location.hostname}:8081/`;

export function apiMethod(name, unautorized) {
    if (name != null && name.length > 0)
        return _.trimEnd(host, '/') + (unautorized ? "/" : "/admin/") + _.trimStart(name, '/');
    else
        return _.trimEnd(host, '/') + "/admin"
}