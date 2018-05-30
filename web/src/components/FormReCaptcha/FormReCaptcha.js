import React from 'react';
import PropTypes from 'prop-types';
import ReCAPTCHA from 'react-google-recaptcha';

import ErrorMessage from 'components/ErrorMessage';

const RE_CAPTCHA_KEY = '6LdTKksUAAAAAJdLyMV7YGi68Kt6tOf0YrW1XUSd';

class ReCaptcha extends React.Component {
    resetRecaptcha() {
        this.reCaptcha.reset();
    }

    render() {
        const { meta: { touched, error, warning }, input: { onChange } } = this.props;
        const showError = !!(touched && (error || warning));

        return (
            <div>
                <ReCAPTCHA ref={cpt => { this.reCaptcha = cpt; }} sitekey={RE_CAPTCHA_KEY} onChange={onChange} />
                {showError && <ErrorMessage>{error || warning}</ErrorMessage>}
            </div>
        );
    }
};

ReCaptcha.propTypes = {
    meta: PropTypes.shape({
        touched: PropTypes.bool,
        error: PropTypes.string,
        warning: PropTypes.string,
    }).isRequired,
    input: PropTypes.shape({
        onChange: PropTypes.func.isRequired,
    }).isRequired,
};

export default ReCaptcha;
