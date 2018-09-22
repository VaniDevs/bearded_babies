import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, TextInput, LongTextInput, SelectInput, ReferenceField, ReferenceInput, ImageField, ImageInput, FormDataConsumer } from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

export const AppointmentIcon = NoteIcon;

export const AppointmentList = (props) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton basePath="/Appointment" />
        </Datagrid>
    </List>
);

const AppointmentTitle = ({ record }) => {
    return <span>Post {record ? `"${record.name}"` : ''}</span>;
};

export const AppointmentCreate = (props) => (
    <Create title="Create an appointment" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth/>
        </SimpleForm>
    </Create>
);