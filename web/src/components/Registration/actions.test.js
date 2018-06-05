import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';
import moxios from 'moxios';
import * as actions from './actions';

const middlewares = [thunk];
const mockStore = configureMockStore(middlewares);

const countriesMock = [
    { Name: 'United States', ISO: 'US' },
    { Name: 'Great Britain', ISO: 'UK' },
    { Name: 'Ukraine', ISO: 'UA' },
];

const seedMock = { seed: 'test seed' };

describe('Registration actions', () => {

    beforeEach(function () {
        moxios.install();
    });

    afterEach(function () {
        moxios.uninstall();
    });

    describe('getCountries', () => {
        it('dispatches GET_COUNTRIES_RESPONSE', () => {
            moxios.wait(() => {
                const request = moxios.requests.mostRecent();
                request.respondWith({
                    status: 200,
                    response: countriesMock,
                });
            });

            const expectedActions = [
                { type: actions.GET_COUNTRIES_RESPONSE, countries: countriesMock },
            ];

            const store = mockStore({ countries: [] });

            return store.dispatch(actions.getCountries()).then(() => {
                expect(store.getActions()).toEqual(expectedActions);
            });
        });

    });

    describe('register', () => {

        it('dispatches ACTIVATE_PROMO_RESPONSE and navigate to congratulations screen, when response is OK', () => {
            moxios.wait(() => {
                const request = moxios.requests.mostRecent();
                request.respondWith({
                    status: 200,
                    response: seedMock,
                });
            });

            const expectedActions = [
                { type: actions.ACTIVATE_PROMO_REQUEST },
                { type: actions.ACTIVATE_PROMO_RESPONSE, payload: seedMock },
                { type: '@@router/CALL_HISTORY_METHOD', payload: { args: ['/thankyou'], method: 'push' } }
            ];

            const store = mockStore({ countries: [] });

            return store.dispatch(actions.register({}, 1, 2)).then(() => {
                expect(store.getActions()).toEqual(expectedActions);
            });
        });
    });
});
