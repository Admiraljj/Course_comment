import axios from 'axios';

export default class CourseService {
    async GetAllCourses(token) {
        console.log("token:" + token);
        const response = await axios.get("http://127.0.0.1:8080/courses", {
            headers: {Authorization: token}
        });
        return response.data.data;
    }

}