import React from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';
import Container from 'components/Container';
import Heading from 'components/Heading';

import icon from './wallet.svg';

const iconHeight = 140;

const Icon = styled.div`
    width: ${iconHeight}px;
    height: ${iconHeight}px;
    transform: translateY(${iconHeight / 3}px);
    margin: auto;
    background-color: ${props => props.theme.colors.white};
    background-image: url(${icon});
    background-position: 50% 50%;
    background-repeat: no-repeat;
    border-radius: 100%;
    
`;

const Seed = styled.div`
    margin-bottom: ${iconHeight / 3}px;
    padding-top: ${iconHeight / 3}px;
    padding-bottom: ${props => props.theme.space[4]}px;
    background: ${props => props.theme.colors.paleGrey};
    text-align: center;
`;

const Text = styled.p`
    font-family: ${props => props.theme.fontLight};
    font-size: ${props => props.theme.fontSizes[1]}px;
`;

class Congratulations extends React.Component {
    render() {
        const { seed } = this.props;

        return (
            <Container>
                <Heading as="h3" fontSize={3} textAlign="center" my={0}>
                    Congratulations!
                </Heading>
                <Icon />
                <Seed>
                    <Text>{seed}</Text>
                </Seed>
            </Container>
        );
    }
}

export default connect(state => ({ seed: state.registration.seed }))(Congratulations);
