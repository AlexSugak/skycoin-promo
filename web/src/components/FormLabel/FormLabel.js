import React from 'react';
import styled from 'styled-components';

const FormLabel = styled.label`
    color: ${props => props.theme.colors.grey};
    display: block;
    margin-bottom: ${ props => props.theme.space[2]}px;
    font-size: ${ props => props.theme.fontSizes[0]}px;
`;

export default FormLabel;
