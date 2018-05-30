import styled from 'styled-components';
import { space, width, fontSize } from 'styled-system';

const Button = styled.button`
    width: ${props => props.block ? '100%': 'auto'};
    padding: ${props => props.theme.space[3]}px ${props => props.theme.space[7]}px;
    border-radius: ${props => props.theme.radius[1]}px;
    border-width: 0;
    background: ${props => props.theme.colors.primary};
    font-size: ${props => props.theme.fontSizes[2]}px;
    font-family: ${props => props.theme.fontLight};
    line-height: 22px;
    color: ${props => props.theme.colors.white};

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
