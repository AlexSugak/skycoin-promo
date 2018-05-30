import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import { routerReducer } from 'react-router-redux';
import registration from 'components/Registration/reducers';

const rootReducer = combineReducers({
    form: formReducer,
    routing: routerReducer,
    registration
});

export default rootReducer;
