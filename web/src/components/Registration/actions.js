import { push } from 'react-router-redux';
import { activatePromo, getCountriesList } from 'api';

export const GET_COUNTRIES_RESPONSE = 'GET_COUNTRIES_RESPONSE';
export const ACTIVATE_PROMO_REQUEST = 'ACTIVATE_PROMO_REQUEST';
export const ACTIVATE_PROMO_RESPONSE = 'ACTIVATE_PROMO_RESPONSE';

export const register = (user, promoId, promoCode) =>
    async dispatch => {
        dispatch({ type: ACTIVATE_PROMO_REQUEST });

        try {
            const response = await activatePromo(user, promoId, promoCode);
            dispatch({ type: ACTIVATE_PROMO_RESPONSE, payload: response.data });
            dispatch(push('/thankyou'));

        } catch (e) {
            if (e.response.status === 400) {
                const errors = e.response.data;
                let formErrors = {};

                if (typeof errors === 'string') {
                    formErrors._error = errors;
                } else {
                    Object.values(errors).map(k => formErrors[k.key] = k.message);
                }

                return Promise.reject(formErrors);
            }
        }
};

export const getCountries = () =>
    async dispatch => {
        try {
            const countries = await getCountriesList().catch(e => { });

            dispatch({
                type: GET_COUNTRIES_RESPONSE,
                countries: countries.data,
            });
        } catch (e) {
            console.log(e.message);
        }
    };