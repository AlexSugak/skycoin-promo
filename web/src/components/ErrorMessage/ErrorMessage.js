import styled from 'styled-components';

const ErrorMessage = styled.p`
    margin-top: ${props => props.theme.space[1]}px;
    margin-bottom: ${props => props.theme.space[1]}px;
    font-size: ${props => props.theme.fontSizes[0]}px;
    color: ${props => props.theme.colors.danger};
`;

export default ErrorMessage;
