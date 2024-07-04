import axios from 'axios';

export default class CommentService {
    async getCommentsByCourseId(id) {
        const response = await axios.get(`http://127.0.0.1:8080/comment/${id}`);
        return response.data.data;
    }

    async addComment(courseId, content, token) {
        try {
            const response = await axios.post(
                "http://127.0.0.1:8080/comment/add",
                {
                    course_id: parseInt(courseId),
                    comment_text: content
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
