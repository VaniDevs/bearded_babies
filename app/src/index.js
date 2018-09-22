import React from 'react';
import { render } from 'react-dom';
import { fetchUtils, Admin, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import authProvider from './authProvider';
import { apiMethod } from './api';

import { AgenciesList, AgenciesEdit, AgenciesCreate, AgenciesIcon } from './agencies';

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    const token = localStorage.getItem('token');
    options.headers.set('Authorization', `Bearer ${token}`);
    return fetchUtils.fetchJson(url, options);
};

const dataProvider = simpleRestProvider(apiMethod(), httpClient);

render(
    <Admin dataProvider={dataProvider} authProvider={authProvider}>
        <Resource name="agencies" list={AgenciesList} edit={AgenciesEdit} create={AgenciesCreate} icon={AgenciesIcon} />
    </Admin>,
    document.getElementById('root')
);