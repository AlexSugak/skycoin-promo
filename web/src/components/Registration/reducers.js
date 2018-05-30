import { GET_COUNTRIES_RESPONSE, ACTIVATE_PROMO_RESPONSE } from './actions';

export const initialState = {
    countries: [],
};

export default (state = initialState, action) => {
    switch (action.type) {
        case GET_COUNTRIES_RESPONSE:
            return { ...state, countries: action.countries.map(c => ({ text: c.Name, value: c.ISO })) };
        case ACTIVATE_PROMO_RESPONSE:
            return { ...state, seed: action.payload.seed };
        default:
            return state;
    }
};
