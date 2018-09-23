import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, DateField, DateInput, TextInput, SelectInput, ReferenceField, ReferenceArrayInput, SelectArrayInput, ReferenceInput, FormDataConsumer } from 'react-admin';
import ReplyAll from '@material-ui/icons/ReplyAll';

const gearStatus = [
    {id: 0, name: "Not needed"},
    {id: 1, name: "Requested"},
    {id: 2, name: "Not available"}
];

export const ReferralsIcon = ReplyAll;

export const ReferralsList = ({ permissions, ...props }) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <ReferenceField label="Client" source="clientId" reference="clients">
                <TextField source="name" />
            </ReferenceField>
            <DateField source="appointment_1"/>
            <DateField source="appointment_2"/>
            <EditButton />
        </Datagrid>
    </List>
);

const ReferralsTitle = ({ record }) => {
    return <span>Referral {record ? `"${record.id}"` : ''}</span>;
};

export const ReferralsEdit = ({ permissions, ...props }) => (
    <Edit title={<ReferralsTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <ReferenceInput label="Client" fullWidth source="clientId" reference="clients">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <ReferenceArrayInput label="Requested gears" fullWidth source="requested" defaultValue={[]} reference="gears">
                <SelectArrayInput optionText="name"/>
            </ReferenceArrayInput>
            <ReferenceArrayInput label="Unavailable gears" fullWidth source="unavailable" defaultValue={[]} reference="gears">
                <SelectArrayInput optionText="name"/>
            </ReferenceArrayInput>
        </SimpleForm>
    </Edit>
);

export const ReferralsCreate = ({ permissions, ...props }) => (
    <Create title="Create a referral" {...props}>
        <SimpleForm>
            <ReferenceInput label="Client" fullWidth source="clientId" reference="clients">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <ReferenceArrayInput label="Requested gears" fullWidth source="requested" defaultValue={[]} reference="gears">
                <SelectArrayInput optionText="name"/>
            </ReferenceArrayInput>
            <ReferenceArrayInput label="Unavailable gears" fullWidth source="unavailable" defaultValue={[]} reference="gears">
                <SelectArrayInput optionText="name"/>
            </ReferenceArrayInput>
        </SimpleForm>
    </Create>
);