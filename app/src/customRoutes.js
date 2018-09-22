import React from 'react';
import { Route } from 'react-router-dom';
import ScheduleAppointment from './ScheduleAppointment';

export default [
    <Route exact path="/schedule-appointment" component={ScheduleAppointment} noLayout />,
];