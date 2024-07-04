import React, { useState, useEffect } from 'react';
import { List, Card, Spin, Pagination, Input, Select } from 'antd';
import CourseService from '../../src/service/CourseService';
import { useRouter } from 'next/router';

const { Option } = Select;
const courseService = new CourseService();

const Courses = () => {
    const [courses, setCourses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [total, setTotal] = useState(0);
    const [currentPage, setCurrentPage] = useState(1);
    const [pageSize] = useState(6); // 每页展示5个课程
    const [searchText, setSearchText] = useState('');
    const [selectedCourseType, setSelectedCourseType] = useState('');
    const router = useRouter();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            fetchCourses(token, currentPage, pageSize);
        } else {
            setLoading(false);
        }
    }, [currentPage]);

    const fetchCourses = async (token, page, pageSize) => {
        try {
            const allCourses = await courseService.GetAllCourses(token);
            const filteredCourses = filterCourses(allCourses, searchText, selectedCourseType);
            setTotal(filteredCourses.length);
            const start = (page - 1) * pageSize;
            const end = start + pageSize;
            setCourses(filteredCourses.slice(start, end));
            setLoading(false);
        } catch (error) {
            console.error('获取课程信息出错:', error);
            setLoading(false);
        }
    };

    const filterCourses = (courses, searchText, courseType) => {
        return courses.filter(course =>
            (course.teacher_name.includes(searchText) ||
                course.course_name.includes(searchText) ||
                course.credits.toString().includes(searchText) ||
                course.course_type.includes(searchText)) &&
            (courseType === '' || courseType === '全部' || course.course_type === courseType)
        );
    };

    const handleCourseClick = (id) => {
        router.push(`/course/${id}`);
    };

    const handlePageChange = (page) => {
        setCurrentPage(page);
    };

    const handleSearch = (value) => {
        setSearchText(value);
        setCurrentPage(1); // 重置为第一页
        const token = localStorage.getItem('token');
        if (token) {
            fetchCourses(token, 1, pageSize);
        }
    };

    const handleCourseTypeChange = (value) => {
        setSelectedCourseType(value);
        setCurrentPage(1); // 重置为第一页
        const token = localStorage.getItem('token');
        if (token) {
            fetchCourses(token, 1, pageSize);
        }
    };

    if (loading) {
        return (
            <div className="flex justify-center items-center h-screen">
                <Spin size="large" />
            </div>
        );
    }

    return (
        <div className="flex justify-center items-center min-h-screen py-10">
            <div className="w-full max-w-4xl">
                <div className="flex mb-4">
                    <Input.Search
                        placeholder="输入教师名、课程名、学分或课程类型进行搜索"
                        allowClear
                        onSearch={handleSearch}
                        className="mr-4"
                    />
                    <Select
                        placeholder="选择课程类型"
                        allowClear
                        onChange={handleCourseTypeChange}
                        className="w-48"
                    >
                        <Option value="全部">全部</Option>
                        <Option value="专业课程">专业课程</Option>
                        <Option value="通识大类">通识大类</Option>
                        <Option value="创新思维与创业实践">创新思维与创业实践</Option>
                        <Option value="人文经典与文化传承">人文经典与文化传承</Option>
                        <Option value="艺术修养与审美体验">艺术修养与审美体验</Option>
                        <Option value="全球视野与文明对话">全球视野与文明对话</Option>
                        <Option value="科学探索与持续发展">科学探索与持续发展</Option>
                    </Select>
                </div>
                <List
                    grid={{ gutter: 16, column: 2 }} // 设置每行展示2个课程
                    dataSource={courses}
                    renderItem={course => (
                        <List.Item>
                            <Card
                                hoverable
                                className="shadow-lg rounded-lg"
                                onClick={() => handleCourseClick(course.ID)}
                            >
                                <div className="text-center">
                                    <h3 className="text-xl font-bold mb-2">{course.course_name}</h3>
                                    <p className="mb-1">学分: {course.credits}</p>
                                    <p className="mb-1">类型: {course.course_type}</p>
                                    <p className="mb-1">教师: {course.teacher_name}</p>
                                </div>
                            </Card>
                        </List.Item>
                    )}
                />
                <div className="flex justify-center mt-4">
                    <Pagination
                        current={currentPage}
                        pageSize={pageSize}
                        total={total}
                        onChange={handlePageChange}
                        className="text-center"
                    />
                </div>
            </div>
        </div>
    );
};

export default Courses;
