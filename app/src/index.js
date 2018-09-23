import React from 'react';
import { render } from 'react-dom';
import { fetchUtils, Admin, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import authProvider from './authProvider';
import { apiMethod } from './api';
import customRoutes from './customRoutes';
import createHistory from 'history/createBrowserHistory';

import { AgenciesList, AgenciesEdit, AgenciesCreate, AgenciesIcon } from './agencies';
import {ClientsList, ClientsEdit, ClientsCreate, ClientsIcon} from './clients';
import {GearList, GearEdit, GearCreate, GearIcon} from './gears';
import {ReferralsList, ReferralsEdit, ReferralsCreate, ReferralsIcon} from './referrals';
import {AppointmentsEdit, AppointmentsIcon} from './appointments';

const history = createHistory();

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
    <Admin customRoutes={customRoutes} dataProvider={dataProvider} authProvider={authProvider} history={history}>
        {permissions => [
        permissions === 'admin' ? <Resource name="agencies" options={{label: "Agencies"}} list={AgenciesList} edit={AgenciesEdit} create={AgenciesCreate} icon={AgenciesIcon} /> : null,
        <Resource name="clients" options={{label: "Clients"}} list={ClientsList} edit={ClientsEdit} create={ClientsCreate} icon={ClientsIcon} />,
        permissions === 'admin' ? <Resource name="gears" options={{label: "Gears"}} list={GearList} edit={GearEdit} create={GearCreate} icon={GearIcon} /> : <Resource name="gears" />,
        <Resource name="referrals" options={{label: "Referrals"}} list={ReferralsList} edit={ReferralsEdit} create={permissions === 'agent' ? ReferralsCreate : null} icon={ReferralsIcon} />,
        ]}
    </Admin>,
    document.getElementById('root')
);