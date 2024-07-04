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

    async DeleteCourseById(id, token) {
        const response = await axios.get(`http://127.0.0.1:8080/course/delete/${id}`, {
            headers: { Authorization: token }
        });
        return response.data;
        }

    async AddCourse(course_name, credits, course_type, teacher_name, token) {
        try {
            const response = await axios.post("http://127.0.0.1:8080/course/add", {
                course_name,
                credits: parseInt(credits),
                course_type,
                teacher_name
            }, {
                headers: { Authorization: token }
            });
            if (response.data.code === 0) {
                return { success: true, data: response.data.data };
            } else {
                return { success: false, message: response.data.message };
            }
        } catch (error) {
            return { success: false, message: error.response.data.message };
        }
    }

}
