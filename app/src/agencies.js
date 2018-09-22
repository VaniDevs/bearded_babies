import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, TextInput, LongTextInput, SelectInput, ReferenceField, ReferenceInput, ImageField, ImageInput, FormDataConsumer } from 'react-admin';
import { required } from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

export const AgenciesIcon = NoteIcon;

export const AgenciesList = (props) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <TextField source="phone" />
            <TextField source="email" />
            <EditButton basePath="/agencies" />
        </Datagrid>
    </List>
);

const AgenciesTitle = ({ record }) => {
    return <span>Agency <b>{record ? `"${record.name}"` : ''}</b></span>;
};

export const AgenciesEdit = (props) => (
    <Edit title={<AgenciesTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth validate={required()} />
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth />
            <TextInput source="email" fullWidth />
            <TextInput source="contact" fullWidth />
        </SimpleForm>
    </Edit>
);

export const AgenciesCreate = (props) => (
    <Create title="Create an agency" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth />
            <TextInput source="email" fullWidth />
            <TextInput source="contact" fullWidth />
        </SimpleForm>
    </Create>
);