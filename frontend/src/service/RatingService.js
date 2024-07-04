import axios from 'axios';

export default class RatingService {
    async getRatingByCourseId(id) {
        const response = await axios.get(`http://127.0.0.1:8080/rating/${id}`);
        return response.data.data;
    }

    async addOrAlterRating(courseId, rating, token) {
        try {
            const response = await axios.post(
                "http://127.0.0.1:8080/rating/add",
                {
                    course_id: parseInt(courseId),
                    rating: rating
                },
                {
                    headers: { Authorization: token }
                }
            );
            return response.data; // 如果你需要返回响应数据
        } catch (error) {
            return error.response.data
        }
    }

}
