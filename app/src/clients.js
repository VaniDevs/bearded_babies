import React from 'react';
import {
    List,
    Datagrid,
    Edit,
    Create,
    SimpleForm,
    TextField,
    EditButton,
    DisabledInput,
    TextInput,
    LongTextInput,
    SelectInput,
    DateField,
    DateInput,
    ReferenceField,
    ReferenceInput,
    ImageField,
    ImageInput,
    BooleanInput,
    FormDataConsumer,
    required
} from 'react-admin';
import NoteIcon from '@material-ui/icons/Note';

export const ClientsIcon = NoteIcon;

export const ClientsList = ({ permissions, ...props }) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <TextField source="phone" />
            <TextField source="email" />
            {permissions === 'admin' ?
                <ReferenceField label="Status" source="status_id" reference="client_statuses" linkType={false}>
                    <TextField source="name" />
                </ReferenceField>
                : null}
            <EditButton basePath="/clients" />
        </Datagrid>
    </List>
);

const ClientsTitle = ({ record }) => {
    return <span>Client <b>{record ? `"${record.name}"` : ''}</b></span>;
};

export const ClientsEdit = ({ permissions, ...props }) => (
    <Edit title={<ClientsTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth validate={required()}/>
            <DateInput source="DOB" fullWidth validate={required()}/>
            <DateInput source="childDOB" fullWidth validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextField source="phone" />
            <TextField source="email" />
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agency_id" reference="agencies">
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <ReferenceInput label="Status" source="status_id" reference="client_statuses">
                    <SelectInput optionText="name" /> : null}
                </ReferenceInput>
                : null}
            <BooleanInput source="unemployed" />
            <BooleanInput source="newcomer" />
            <BooleanInput source="homeless" />
            <BooleanInput source="special_needs" />
        </SimpleForm>
    </Edit>
);

export const ClientsCreate = ({ permissions, ...props }) => (
    <Create title="Create a client" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth validate={required()}/>
            <DateInput source="DOB" fullWidth validate={required()}/>
            <DateInput source="childDOB" fullWidth validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextField source="phone" />
            <TextField source="email" />
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agency_id" reference="agencies">
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <ReferenceInput label="Status" source="status_id" reference="client_statuses">
                    <SelectInput optionText="name" /> : null}
                </ReferenceInput>
                : null}
            <BooleanInput source="unemployed" />
            <BooleanInput source="newcomer" />
            <BooleanInput source="homeless" />
            <BooleanInput source="special_needs" />
        </SimpleForm>
    </Create>
);