import React from 'react';
import styled from 'styled-components';
import FormGroup from 'components/FormGroup';

const ControlInput = styled.input`
    width: 100%;
    height: ${props => props.theme.controlHeight}px;
    padding: ${props => props.theme.space[1]}px ${props => props.theme.space[2]}px;
    border-width: 1px;
    border-style: solid;
    border-color: ${props => props.theme.colors.lightGrey};
    border-radius: ${props => props.theme.radius[0]}px;
    background: transparent;
    font-family: ${props => props.theme.fontLight};
    font-size: ${props => props.theme.fontSizes[1]}px;
    line-height: ${props => props.theme.lineHeights[2]};
    color: ${props => props.theme.colors.darkGrey};
    
    &:focus {
        outline: none;
        border: 1px solid ${props => props.theme.colors.primary};
    }
`;

const FormControl = (props) => {
    const { label, placeholder, type, maxLength, input, meta: { error, warning, touched } } = props;
    const showError = !!(touched && (error || warning));

    return (
        <FormGroup label={label} placeholder={placeholder} name={input.name}>
            <ControlInput value={input.value} name={input.name} type={type} placeholder={placeholder} maxLength={maxLength} onChange={input.onChange} />
        </FormGroup>
    );
}

FormControl.defaultProps = {
    placeholder: '',
};

export default FormControl;
