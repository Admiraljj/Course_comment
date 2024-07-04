import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import { Card, List, Pagination, message, Rate } from 'antd';
import 'tailwindcss/tailwind.css';
import CourseService from '../../src/service/CourseService';
import CommentService from '../../src/service/CommentService';
import RatingService from '../../src/service/RatingService';

const CourseDetail = () => {
    const router = useRouter();
    const { id } = router.query;
    const [courseInfo, setCourseInfo] = useState(null);
    const [comments, setComments] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [pageSize, setPageSize] = useState(3);
    const [newComment, setNewComment] = useState('');
    const [rating, setRating] = useState(0);
    const [averageRating, setAverageRating] = useState(0);
    const desc = ['非常差', '孬', '一般', '好', '棒极啦'];

    useEffect(() => {
        if (id) {
            fetchCourseInfo(id);
            fetchComments(id);
            fetchAverageRating(id);
        }
    }, [id]);

    const fetchCourseInfo = async (id) => {
        const courseService = new CourseService();
        const data = await courseService.GetCourseInfoById(id);
        setCourseInfo(data);
    };

    const fetchComments = async (id) => {
        const commentService = new CommentService();
        const data = await commentService.getCommentsByCourseId(id);
        setComments(data);
    };

    const fetchAverageRating = async (id) => {
        const ratingService = new RatingService();
        const data = await ratingService.getRatingByCourseId(id);
        const ratings = data.map(r => r.rating);
        const average = ratings.reduce((a, b) => a + b, 0) / ratings.length;
        setAverageRating(average);
        setRating(average);
    };

    const handlePageChange = (page, pageSize) => {
        setCurrentPage(page);
        setPageSize(pageSize);
    };

    const handleAddComment = async () => {
        if (!newComment.trim()) {
            message.error('评论内容不能为空');
            return;
        }

        try {
            const commentService = new CommentService();
            const token = localStorage.getItem('token');
            const response = await commentService.addComment(id, newComment, token);
            if (response.code === 3) {
                message.error(response.message);
            } else {
                message.success('评论提交成功');
                fetchComments(id); // 重新获取评论
                setNewComment(''); // 清空输入框
            }
        } catch (error) {
            message.error('评论提交失败');
        }
    };

    const handleRateChange = async (value) => {
        setRating(value);
        try {
            const ratingService = new RatingService();
            const token = localStorage.getItem('token');
            const response = await ratingService.addOrAlterRating(id, value, token);
            if (response.code === 3) {
                message.error(response.message);
            } else {
                message.success('评分提交成功');
                fetchAverageRating(id); // 重新获取综合评分
            }
        } catch (error) {
            message.error('评分提交失败');
        }
    };

    const handleBackToHome = () => {
        router.push('/'); // 返回主页路径
    };

    return (
        <div className="p-4 max-w-3xl mx-auto">
            <h1 className="text-3xl font-bold mb-4 text-center">课程详情</h1>
            {courseInfo && (
                <Card className="mb-4 shadow-lg rounded-lg border border-gray-200">
                    <div className="flex justify-between items-center">
                        <div>
                            <p><strong>课程名称：</strong>{courseInfo.course_name}</p>
                            <p><strong>学分：</strong>{courseInfo.credits}</p>
                            <p><strong>课程类型：</strong>{courseInfo.course_type}</p>
                            <p><strong>教师名称：</strong>{courseInfo.teacher_name}</p>
                        </div>
                        <div className="flex flex-col items-center">
                            <Rate tooltips={desc} onChange={handleRateChange} value={rating} />
                            {rating ? <span className="ant-rate-text">{desc[rating - 1]}</span> : ''}
                            <p>当前综合评分：{averageRating.toFixed(1)}</p>
                        </div>
                    </div>
                </Card>
            )}
            <hr className="my-4"/>

            <List
                dataSource={comments.slice((currentPage - 1) * pageSize, currentPage * pageSize)}
                renderItem={comment => (
                    <List.Item className="mb-4">
                        <Card className="w-full shadow-lg rounded-lg border border-gray-200">
                            <div className="flex items-start">
                                <img
                                    src="/avatar.jpg"
                                    alt="User Avatar"
                                    className="h-12 w-12 rounded-full border-2 border-white mr-4 cursor-pointer"
                                />
                                <div className="flex-1">
                                    <p><strong>{comment.user_name}</strong></p>
                                    <p className="text-gray-800">{comment.comment_text}</p>
                                    <p className="text-gray-400 text-xs text-right">{new Date(comment.comment_date).toLocaleString()}</p>
                                </div>
                            </div>
                        </Card>
                    </List.Item>
                )}
            />
            <Pagination
                className="mt-4 text-center"
                current={currentPage}
                pageSize={pageSize}
                total={comments.length}
                onChange={handlePageChange}
            />
            <div className="mt-4 mb-4">
                <textarea
                    className="w-full p-2 border border-gray-200 rounded-lg"
                    placeholder="请输入评论内容"
                    value={newComment}
                    onChange={(e) => setNewComment(e.target.value)}
                />
                <button
                    className="mt-2 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg"
                    onClick={handleAddComment}
                >提交评论
                </button>
                <button
                    className="ml-2 bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-lg"
                    onClick={handleBackToHome}
                >
                    返回主页
                </button>
            </div>
        </div>
    );
};

export default CourseDetail;
