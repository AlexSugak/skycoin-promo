import React from 'react';
import PropTypes from 'prop-types';
import { Field, reduxForm, Form } from 'redux-form';
import { createTextMask } from 'redux-form-input-masks';
import { Flex, Box } from 'grid-styled';
import Button from 'components/Button';
import FormControl from 'components/FormControl';
import FormReCaptcha from 'components/FormReCaptcha';
import Heading from 'components/Heading';
import { required, email } from 'validation';

const phoneMask = createTextMask({
    pattern: '(999) 99 999-99-99',
});


class RegistrationForm extends React.Component {
    render() {
        const { handleSubmit, pristine, submitting } = this.props;

        return (
            <Form onSubmit={handleSubmit} noValidate>
                <Heading mb={5} fontSize={2}>Name</Heading>
                <Flex mx={-4} mb={6} flexWrap="wrap">
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="firstName" label="First Name" component={FormControl} validate={[required]} />
                    </Box>
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="lastName" label="Last Name" component={FormControl} validate={[required]} />
                    </Box>
                </Flex>
                <Heading mb={5} fontSize={2}>Contact Info</Heading>
                <Flex mx={-4} flexWrap="wrap">
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="email" name="email" label="Email" component={FormControl} validate={[required, email]} />
                    </Box>
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="tel" name="mobile" label="Mobile Phone" component={FormControl} validate={[required]} {...phoneMask} />
                    </Box>
                </Flex>
                <Field type="text" name="addressLine1" label="Address line 1" component={FormControl} validate={[required]} />
                <Field type="text" name="addressLine2" label="Address line 2" component={FormControl} validate={[required]} />
                <Flex mx={-4} flexWrap="wrap">
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="city" label="City" component={FormControl} validate={[required]} />
                    </Box>
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="state" label="State/region" component={FormControl} validate={[required]} />
                    </Box>
                </Flex>
                <Flex mx={-4} flexWrap="wrap">
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="postcode" label="Postcode" component={FormControl} validate={[required]} />
                    </Box>
                    <Box width={[1, 1, 1 / 2]} px={4}>
                        <Field type="text" name="countryCode" label="Country" component={FormControl} validate={[required]} />
                    </Box>
                </Flex>
                <Flex mt={6} flexWrap="wrap" alignItems="center" justifyContent="space-between">
                        <Field name="recaptcha" component={FormReCaptcha} validate={[required]} withRef ref={r => { this.recaptchaField = r }} />
                        <Button disabled={pristine || submitting} primary>Sign Up</Button>
                </Flex>
            </Form>
        )
    }
}

RegistrationForm.propTypes = {
    handleSubmit: PropTypes.func.isRequired,
    pristine: PropTypes.bool.isRequired,
    submitting: PropTypes.bool.isRequired,
};

export default reduxForm({
    form: 'registrationForm'
})(RegistrationForm);
