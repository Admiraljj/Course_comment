import axios from 'axios';

export default class UserService {
    async signIn(username, password) {
        try {
            console.log("username:" + username + " password:" + password);
            const response = await axios.post('http://127.0.0.1:8080/user/login', {
                username: username,
                password: password
            });
            // 将 token 存入浏览器的 localStorage 中
            console.log("response.data.data.token:" + response.data.data.token)
            return response.data;
        }catch (error) {
            return error.response.data
        }
    }

    async signUp(username, email, password) {
        try {
            const response = await axios.post('http://127.0.0.1:8080/user/register',
                {
                    username: username,
                    email: email,
                    password: password
            });
            console.log("response.data:" + response.data.message);
            return response.data
        } catch (error) {
            return error.response.data
        }

    }

    async signOut() {
        // 清除浏览器的 localStorage 中的 token
        localStorage.removeItem('token');
    }

    async GetUserInfoByToken(token) {
        console.log("token:" + token);
        const response = await axios.get("http://127.0.0.1:8080/user/info", {
            headers: {Authorization: token}
        });
        return response.data;
    }
}