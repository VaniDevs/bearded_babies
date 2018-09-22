import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, TextInput, LongTextInput, SelectInput, ReferenceField, ReferenceInput, ImageField, ImageInput, FormDataConsumer } from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

export const ClientsIcon = NoteIcon;

export const ClientsList = (props) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton basePath="/users" />
        </Datagrid>
    </List>
);

const ClientsTitle = ({ record }) => {
    return <span>Post {record ? `"${record.name}"` : ''}</span>;
};

export const ClientsEdit = (props) => (
    <Edit title={<ClientsTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth />
        </SimpleForm>
    </Edit>
);

export const ClientsCreate = (props) => (
    <Create title="Create a user" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth/>
        </SimpleForm>
    </Create>
);