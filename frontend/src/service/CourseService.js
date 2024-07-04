import axios from 'axios';

export default class CourseService {
    async GetAllCourses() {
        const response = await axios.get("http://127.0.0.1:8080/courses");
        return response.data.data;
    }

    async GetCourseInfoById(id) {
        const response = await axios.get(`http://127.0.0.1:8080/course/${id}`);
        return response.data.data;
    }


}
