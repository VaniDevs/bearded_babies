import React from 'react';
import {
    List,
    Datagrid,
    Edit,
    Create,
    SimpleForm,
    SelectField,
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

const clientStatus = [
    {id: 0, name: "Pending"},
    {id: 1, name: "Approved"},
    {id: 2, name: "Declined"}
];

export const ClientsIcon = NoteIcon;

export const ClientsList = ({ permissions, ...props }) => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <TextField source="phone" />
            <TextField source="email" />
            {permissions === 'admin' ?
                <SelectField source="name" choices={clientStatus}/>
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
            <DateInput label="DOB" source="dob" validate={required()}/>
            <DateInput label="Child DOB" source="childdob" validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth/>
            <TextInput source="email" fullWidth/>
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agency_id" reference="agencies" filter={{role: 2}} validate={required()}>
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <SelectInput label="Status" optionText="name" choices={clientStatus} validate={required()}/>
                : null
            }
            <BooleanInput source="unemployed" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="newcomer" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="homeless" format={v => !!v} parse={v => +v}/>
            <BooleanInput label="Special needs" source="special_needs" format={v => !!v} parse={v => +v}/>
        </SimpleForm>
    </Edit>
);

export const ClientsCreate = ({ permissions, ...props }) => (
    <Create title="Create a client" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth validate={required()}/>
            <DateInput label="DOB" source="dob" validate={required()}/>
            <DateInput label="Child DOB" source="childdob" validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth/>
            <TextInput source="email" fullWidth/>
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agency_id" reference="agencies" filter={{role: 2}} validate={required()}>
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <SelectInput label="Status" optionText="name" choices={clientStatus} validate={required()}/>
                : null
            }
            <BooleanInput source="unemployed" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="newcomer" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="homeless" format={v => !!v} parse={v => +v}/>
            <BooleanInput label="Special needs" source="special_needs" format={v => !!v} parse={v => +v}/>
        </SimpleForm>
    </Create>
);