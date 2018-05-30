import React from 'react';
import PropTypes from 'prop-types';
import { Switch, Route } from 'react-router-dom';

import Registration from './Registration';
import Congratulations from './Congratulations';

const Routes = ({ match }) => {
    return (
        <Switch>
            <Route path={`/:promoId/:code`} component={Registration} exact />
            <Route path={`/thankyou`} component={Congratulations} exact />
        </Switch>
    );
};

Routes.propTypes = {
    match: PropTypes.shape({
        url: PropTypes.string.isRequired,
    }).isRequired,
};

export default Routes;
