import styled from 'styled-components';
import { space, width, fontSize } from 'styled-system';

const Button = styled.button`
    padding: 8px 26px;
    font-family: ${props => props.theme.fontLight};
    line-height: 22px;
    font-size: ${props => props.theme.fontSizes[1]}px;
    border-width: 1px;
    width: ${props => props.block ? '100%': 'auto'};

    &:focus {
        outline: none;
    }
    &:hover {
        cursor: pointer;
    }
    &:disabled {
        opacity: 0.5;
    }
    
    ${width}
    ${space}
    ${fontSize}
`;

export default Button;
