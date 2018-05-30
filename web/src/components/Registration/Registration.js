import React from 'react';
import { SubmissionError } from 'redux-form';
import styled from 'styled-components';
import { connect } from 'react-redux';
import Container from 'components/Container';
import Heading from 'components/Heading';

import RegistrationForm from './RegistrationForm';
import { register, getCountries } from './actions';

const StyledHeading = styled(Heading)`
    border-bottom: 2px solid ${props => props.theme.colors.lightGrey};
`;

class Registration extends React.Component {
    componentWillMount() {
        this.props.getCountries();
    }

    onSubmit = (user) => {
        const { match: { params }, activate } = this.props;

        return activate(user, params.promoId, params.code)
            .catch(error => {
                throw new SubmissionError(error);
            });
    };

    render() {
        return (
            <Container>
                <Heading as="h3" fontSize={3} textAlign="center" my={4}>
                    Hello
                </Heading>
                <StyledHeading as="h1" heavy fontSize={2} textAlign="center" pb={4} mb={8}>
                    Register to receive 1 free SKY.  If you have not done so already, please download the wallet from <a href="https://www.skycoin.net/downloads/">here</a>
                </StyledHeading>

                <RegistrationForm onSubmit={this.onSubmit} countries={this.props.countries} />
            </Container>
        );
    }
}

export default connect(state => ({ countries: state.registration.countries }), { activate: register, getCountries })(Registration);
