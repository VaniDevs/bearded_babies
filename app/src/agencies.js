import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, TextInput, LongTextInput, SelectInput, ReferenceField, ReferenceInput, ImageField, ImageInput, FormDataConsumer } from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

export const AgenciesIcon = NoteIcon;

export const AgenciesList = (props) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton basePath="/agencies" />
        </Datagrid>
    </List>
);

const AgenciesTitle = ({ record }) => {
    return <span>Post {record ? `"${record.name}"` : ''}</span>;
};

export const AgenciesEdit = (props) => (
    <Edit title={<AgenciesTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth />
        </SimpleForm>
    </Edit>
);

export const AgenciesCreate = (props) => (
    <Create title="Create an agency" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth/>
        </SimpleForm>
    </Create>
);