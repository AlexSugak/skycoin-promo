import React from 'react';
import PropTypes from 'prop-types';

import FormLabel from 'components/FormLabel';
import ErrorMessage from 'components/ErrorMessage';

import styled from 'styled-components';

const Wrapper = styled.div`
    margin-top: ${props => props.theme.space[1]}px;
    margin-bottom: ${props => props.theme.space[4]}px;
`;

const FormGroup = ({ name, label, isRequired, description, showError, error, children }) => {
    return (
        <Wrapper label={label}>
            <FormLabel for={name} isRequired={isRequired}>{label}</FormLabel>
            {children}
            {showError && <ErrorMessage>{error}</ErrorMessage>}
        </Wrapper>
    );
}

FormGroup.propTypes = {
    name: PropTypes.string.isRequired,
    label: PropTypes.oneOfType([PropTypes.string, PropTypes.object]),
    isRequired: PropTypes.bool,
    description: PropTypes.string,
    showError: PropTypes.bool,
    error: PropTypes.string
};

export default FormGroup;
