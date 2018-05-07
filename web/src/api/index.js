import axios from 'axios';

const headers = {
    headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
    }
};

export const activatePromo = user =>
    axios.post(`/promo/activate`, JSON.stringify(user), headers);

