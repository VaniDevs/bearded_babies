import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, DateField, DateInput, TextInput, SelectInput, ReferenceField, ReferenceArrayInput, SelectArrayInput, ReferenceInput, FormDataConsumer } from 'react-admin';
import ScheduleIcon from '@material-ui/icons/Schedule';

export const AppointmentsIcon = ScheduleIcon;

export const AppointmentsEdit = ({ permissions, ...props }) => (
    <Edit title="Schedule an appointment..." {...props}>
        <SimpleForm>
            <DisabledInput source="id" />

        </SimpleForm>
    </Edit>
);
