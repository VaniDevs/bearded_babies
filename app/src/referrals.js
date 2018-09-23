import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, TextField, EditButton, DisabledInput, DateField, DateInput, TextInput, LongTextInput, SelectInput, ReferenceField, ReferenceInput, ImageField, ImageInput, FormDataConsumer } from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

const gearStatus = [
    {id: 0, name: "Not needed"},
    {id: 1, name: "Needed"},
    {id: 2, name: "Not available"}
];


export const ReferralsIcon = NoteIcon;

export const ReferralsList = ({ permissions, ...props }) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <ReferenceField label="Client" source="client_id" reference="clients">
                <TextField source="name" />
            </ReferenceField>
            <DateField source="appointment1"/>
            <DateField source="appointment2"/>
            <EditButton basePath="/referrals" />
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
            <ReferenceInput label="Client" source="client_id" reference="clients">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <DateInput source="appointment1_date"/>
            <FormDataConsumer>
                {({ formData }) =>
                    <ReferenceInput label="Appointment1 Time" source="appointment1" reference="appointments" filter={{date: formData.appointment1_date}}>
                        <SelectInput optionText="date" />
                    </ReferenceInput>
                }
            </FormDataConsumer>
            <DateInput source="appointment2_date"/>
            <FormDataConsumer>
                {({ formData }) =>
                    <ReferenceInput label="Appointment2 Time" source="appointment2" reference="appointments" filter={{date: formData.appointment2_date}}>
                        <SelectInput optionText="date" />
                    </ReferenceInput>
                }
            </FormDataConsumer>
        </SimpleForm>
    </Edit>
);

export const ReferralsCreate = ({ permissions, ...props }) => (
    <Create title="Create a referral" {...props}>
        <SimpleForm>
            <ReferenceInput label="Client" source="client_id" reference="clients">
                <SelectInput optionText="name" />
            </ReferenceInput>
        </SimpleForm>
    </Create>
);