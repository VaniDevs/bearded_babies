import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, TextInput, SelectInput, ReferenceField, ReferenceInput, FormDataConsumer } from 'react-admin';
import ChildFriendly from '@material-ui/icons/ChildFriendly';

export const GearIcon = ChildFriendly;

export const GearList = (props) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton />
        </Datagrid>
    </List>
);

const GearTitle = ({ record }) => {
    return <span>Gear <b>{record ? `"${record.name}"` : ''}</b></span>;
};

export const GearEdit = (props) => (
    <Edit title={<GearTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth />
        </SimpleForm>
    </Edit>
);

export const GearCreate = (props) => (
    <Create title="Create a gear" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth/>
        </SimpleForm>
    </Create>
);