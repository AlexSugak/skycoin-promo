import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import FormGroup from 'components/FormGroup';
import ErrorMessage from 'components/ErrorMessage';

const SelectWrapper = styled.div`
    position: relative;
    
    &::after {
        content: '';
        position: absolute;
        top: 50%;
        right: ${props => props.theme.space[3]}px;
        width: 0;
        height: 0;
        border-style: solid;
        border-width: 5px 4px 0 4px;
        border-color: ${props => props.theme.colors.grey} transparent transparent transparent;
        transform: translateY(-50%);
    }
`;

const Select = styled.select`
    width: 100%;
    padding: ${props => props.theme.space[1]}px ${props => props.theme.space[2]}px;
    border-width: 1px;
    border-style: solid;
    border-color: ${props => (props.showError ? props.theme.colors.danger : props.theme.colors.lightGrey)};
    border-radius: ${props => props.theme.radius[0]}px;
    background: transparent;
    font-family: ${props => props.theme.fontLight};
    font-size: ${props => props.theme.fontSizes[1]}px;
    line-height: ${props => props.theme.lineHeights[4]};
    color: ${props => props.theme.colors.darkGrey};
    appearance: none;
    
    &::-ms-expand {
        display: none;
    }
    
    &:focus {
        outline: none;
        border: 1px solid ${props => props.theme.colors.primary};
    }
`;

const Dropdown = ({ name, options, onChange, error, input }) =>
    <SelectWrapper>
        <Select name={name} value={input && input.value} onChange={onChange} error={error}>
            <option value="" disabled>Select</option>
            {options.map((item, i) => <option value={item.value} key={i}>{item.text}</option>)}
        </Select>
    </SelectWrapper>;

Dropdown.propTypes = {
    name: PropTypes.string.isRequired,
    onChange: PropTypes.func.isRequired,
    options: PropTypes.arrayOf(PropTypes.shape({
        value: PropTypes.any.isRequired,
        text: PropTypes.string.isRequired,
    })),
};

const FormDropdown = props => {
    const { label, options, description, input: { name, onChange }, meta: { error, warning, touched } } = props;
    const showError = !!(touched && (error || warning));

    return (
        <FormGroup name={name} label={label} description={description}>
            <Dropdown {...props} name={name} onChange={onChange} options={options} />
            {showError && <ErrorMessage>{error || warning}</ErrorMessage>}
        </FormGroup>
    );
};


FormDropdown.propTypes = {
    input: PropTypes.shape({
        onChange: PropTypes.func.isRequired,
        name: PropTypes.string.isRequired,
    }).isRequired,
    meta: PropTypes.shape({
        touched: PropTypes.bool,
        error: PropTypes.string,
        warning: PropTypes.string,
    }).isRequired,
    label: PropTypes.string,
    isRequired: PropTypes.bool,
    options: PropTypes.array,
};

export default FormDropdown;
