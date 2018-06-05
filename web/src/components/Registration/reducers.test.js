import reducer, { initialState } from './reducers';
import { GET_COUNTRIES_RESPONSE, ACTIVATE_PROMO_RESPONSE } from './actions';

const countriesResponse = [
    { Name: 'United States', ISO: 'US' },
    { Name: 'Great Britain', ISO: 'UK' },
    { Name: 'Ukraine', ISO: 'UA' },
];

const countries = [
    { text: 'United States', value: 'US' },
    { text: 'Great Britain', value: 'UK' },
    { text: 'Ukraine', value: 'UA' },
];

describe('Registration reducer', () => {
    describe('countries', () => {

        it('saves countries to display in dropdown', () => {
            const state = reducer(initialState, { type: GET_COUNTRIES_RESPONSE, countries: countriesResponse });
            expect(state.countries).toEqual(countries);
        });

        it('saves seed', () => {
            const state = reducer(initialState, { type: ACTIVATE_PROMO_RESPONSE, payload: { seed: 'test seed' } });
            expect(state.seed).toEqual('test seed');
        });

    });
});
