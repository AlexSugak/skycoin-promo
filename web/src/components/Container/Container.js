import React from 'react';
import styled from 'styled-components';
import { Box } from 'grid-styled';

const Container = styled(Box)`
    ${props => props.width ? '' : `width: ${props.theme.container.width};`}
    max-width: ${props => props.theme.container.maxWidth};
`;

Container.defaultProps = {
    mx: 'auto',
};

export default (props) => (
    <Container {...props}>
        {props.children}
    </Container>
);
