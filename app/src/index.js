import React from 'react';
import { render } from 'react-dom';
import { fetchUtils, Admin, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import authProvider from './authProvider';
import { apiMethod } from './api';
import customRoutes from './customRoutes';

import { AgenciesList, AgenciesEdit, AgenciesCreate, AgenciesIcon } from './agencies';
import { UsersList, UsersEdit, UsersCreate, UsersIcon } from './users';
import {ClientsList, ClientsEdit, ClientsCreate, ClientsIcon} from './clients';
import {GearList, GearEdit, GearCreate, GearIcon} from './gear';
import {ReferralsList, ReferralsEdit, ReferralsCreate, ReferralsIcon} from './referrals';

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
    <Admin customRoutes={customRoutes} dataProvider={dataProvider} authProvider={authProvider}>
        <Resource name="agencies" list={AgenciesList} edit={AgenciesEdit} create={AgenciesCreate} icon={AgenciesIcon} />
        <Resource name="clients" list={ClientsList} edit={ClientsEdit} create={ClientsCreate} icon={ClientsIcon} />
        <Resource name="users" list={UsersList} edit={UsersEdit} create={UsersCreate} icon={UsersIcon} />
        <Resource name="gear" list={GearList} edit={GearEdit} create={GearCreate} icon={GearIcon} />
        <Resource name="Referrals" list={ReferralsList} edit={ReferralsEdit} create={ReferralsCreate} icon={ReferralsIcon} />
    </Admin>,
    document.getElementById('root')
);