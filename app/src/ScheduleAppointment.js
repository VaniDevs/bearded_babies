import React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import { Title } from 'react-admin';

const ScheduleAppointment = ({location, ...props}) => {
    console.log(location);
    console.log(props);
    return <Card>
        <Title title="Schedule appointment" />
        <CardContent>

        </CardContent>
    </Card>
};

export default ScheduleAppointment;