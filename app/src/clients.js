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
    SelectInput,
    DateField,
    DateInput,
    ReferenceField,
    ReferenceInput,
    BooleanInput,
    FormDataConsumer,
    required,
    Filter
} from 'react-admin';
import AccountCircle from '@material-ui/icons/AccountCircle';

const clientStatus = [
    {id: 0, name: "New"},
    {id: 1, name: "Approved"},
    {id: 2, name: "Declined"}
];

export const ClientsIcon = AccountCircle;

const ClientsFilter = (props) => (
    <Filter {...props}>
        <SelectInput label="Status" optionText="name" source="status" choices={clientStatus}/>
    </Filter>
);

export const ClientsList = ({ permissions, ...props }) => (
    <List {...props} filters={permissions === 'admin' ? <ClientsFilter/> : null}>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            {permissions === 'admin' ?
                <ReferenceField label="Agency" source="agencyId" reference="agencies">
                    <TextField source="name" />
                </ReferenceField>
            : null}
            <TextField source="phone" />
            <TextField source="email" />
            {permissions === 'admin' ?
                <SelectField source="status" optionText="name" choices={clientStatus}/>
                : null}
            <EditButton />
        </Datagrid>
    </List>
);

const ClientsTitle = ({ record }) => {
    return <span>Client <b>{record ? `"${record.name}"` : ''}</b></span>;
};

const parseDate = function (v) {
    return v + "T00:00:00Z"
};

export const ClientsEdit = ({ permissions, ...props }) => (
    <Edit title={<ClientsTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="name" fullWidth validate={required()}/>
            <DateInput label="DOB" source="dob" parse={parseDate} validate={required()}/>
            <DateInput label="Child DOB" source="childDob" parse={parseDate} validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth/>
            <TextInput source="email" fullWidth/>
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agencyId" reference="agencies" filter={{role: 2}} validate={required()}>
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <SelectInput label="Status" optionText="name" source="status" choices={clientStatus} validate={required()}/>
                : null
            }
            <BooleanInput source="unemployed" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="newcomer" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="homeless" format={v => !!v} parse={v => +v}/>
            <BooleanInput label="Special needs" source="specialNeeds" format={v => !!v} parse={v => +v}/>
        </SimpleForm>
    </Edit>
);

export const ClientsCreate = ({ permissions, ...props }) => (
    <Create title="Create a client" {...props}>
        <SimpleForm>
            <TextInput source="name" fullWidth validate={required()}/>
            <DateInput label="DOB" source="dob" parse={parseDate} validate={required()}/>
            <DateInput label="Child DOB" source="childDob" parse={parseDate} validate={required()}/>
            <TextInput source="address1" fullWidth validate={required()}/>
            <TextInput source="address2" fullWidth />
            <TextInput source="city" fullWidth validate={required()}/>
            <TextInput source="phone" fullWidth/>
            <TextInput source="email" fullWidth/>
            {permissions === 'admin' ?
                <ReferenceInput label="Agency" source="agencyId" reference="agencies" filter={{role: 2}} validate={required()}>
                    <SelectInput optionText="name" />
                </ReferenceInput>
                : null}
            {permissions === 'admin' ?
                <SelectInput label="Status" optionText="name" source="status" choices={clientStatus} validate={required()}/>
                : null
            }
            <BooleanInput source="unemployed" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="newcomer" format={v => !!v} parse={v => +v}/>
            <BooleanInput source="homeless" format={v => !!v} parse={v => +v}/>
            <BooleanInput label="Special needs" source="specialNeeds" format={v => !!v} parse={v => +v}/>
        </SimpleForm>
    </Create>
);