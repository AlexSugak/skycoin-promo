import { activatePromo } from 'api';

export const ACTIVATE_PROMO_REQUEST = 'ACTIVATE_PROMO_REQUEST';
export const ACTIVATE_PROMO_RESPONSE = 'ACTIVATE_PROMO_RESPONSE';

export const register = (user) =>
    async dispatch => {
        dispatch({ type: ACTIVATE_PROMO_REQUEST });

        await activatePromo(user);
        dispatch({ type: ACTIVATE_PROMO_RESPONSE });
};
