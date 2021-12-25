import axios from 'axios'

class GotestAPI {
    getList() {
        return axios.get('/api/users');
    }

    create(data) {
        return axios.post('/api/users', data)
    }
}
const gotestApi = new GotestAPI()

export default gotestApi;