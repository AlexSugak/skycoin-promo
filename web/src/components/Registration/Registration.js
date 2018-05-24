import React from 'react';
import styled from 'styled-components';
import { connect } from 'react-redux';
import Container from 'components/Container';
import Heading from 'components/Heading';

import RegistrationForm from './RegistrationForm';
import { register } from './actions';

const StyledHeading = styled(Heading)`
    border-bottom: 2px solid ${props => props.theme.colors.lightGrey};
`;

class Registration extends React.PureComponent {
    onSubmit = (user) => {
        this.props.registerUser(user);
    }

    render() {
        return (
            <Container>
                <Heading as="h3" fontSize={3} textAlign="center" my={4}>
                    Hello
                </Heading>
                <StyledHeading as="h1" heavy fontSize={2} textAlign="center" pb={4} mb={8}>
                    Register to receive 1 free SKY.  If you have not done so already, please download the wallet from <a href="https://www.skycoin.net/downloads/">here</a>
                </StyledHeading>

                <RegistrationForm onSubmit={this.onSubmit}  />
            </Container>
        );
    }
}

export default connect(null, { registerUser: register })(Registration);
