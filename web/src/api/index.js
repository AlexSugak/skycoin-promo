import axios from 'axios';

const headers = {
    headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
    }
};

export const activatePromo = (user, promoId, promoCode) =>
    axios.post(`/promo/${promoId}/${promoCode}`, JSON.stringify(user), headers);

export const getCountriesList = () =>
    axios.get(`/geo/countries`, headers);