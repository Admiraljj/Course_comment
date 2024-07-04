import {DownOutlined, PlusOutlined} from '@ant-design/icons';
import {Button, Card, Form, Input, message, Modal, Pagination, Radio, Select, Space, Switch, Table} from 'antd';
import React, { useState, useEffect } from 'react';
import CourseService from '../../src/service/CourseService';

const { Option } = Select;

const App = () => {
    const [data, setData] = useState([]);
    const [filteredData, setFilteredData] = useState([]);
    const [selectedRowKeys, setSelectedRowKeys] = useState([]);
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [pageSize] = useState(6);
    const [searchText, setSearchText] = useState('');
    const [selectedCourseType, setSelectedCourseType] = useState('');
    const [form] = Form.useForm();

    const fetchData = async () => {
        const courseService = new CourseService();
        const result = await courseService.GetAllCourses();
        setData(result);
        setFilteredData(result);
    };

    useEffect(() => {
        fetchData();
    }, []);

    const handleDelete = async (ids) => {
        const courseService = new CourseService();
        const token = localStorage.getItem('token'); // 从localStorage获取token
        await Promise.all(ids.map(id => courseService.DeleteCourseById(id, token)));
        setData(data.filter((item) => !ids.includes(item.ID)));
        setFilteredData(filteredData.filter((item) => !ids.includes(item.ID)));
        setSelectedRowKeys([]); // 清空选中的行
    };

    const handleAddCourse = async (values) => {
        const courseService = new CourseService();
        const token = localStorage.getItem('token'); // 从localStorage获取token
        const result = await courseService.AddCourse(values.course_name, values.credits, values.course_type, values.teacher_name, token);
        if (result.success) {
            message.success('课程添加成功');
            fetchData(); // 重新获取所有课程信息
        } else {
            message.error(result.message);
        }
    };

    const showAddCourseModal = () => {
        setIsModalVisible(true);
    };

    const handleOk = () => {
        form
            .validateFields()
            .then((values) => {
                form.resetFields();
                handleAddCourse(values);
                setIsModalVisible(false);
            })
            .catch((info) => {
                console.log('Validate Failed:', info);
            });
    };

    const handleCancel = () => {
        setIsModalVisible(false);
    };

    const handleSearch = (value) => {
        setSearchText(value);
        applyFilter(value, selectedCourseType);
    };

    const handleCourseTypeChange = (value) => {
        setSelectedCourseType(value);
        applyFilter(searchText, value);
    };

    const applyFilter = (searchText, courseType) => {
        const filteredCourses = data.filter(course =>
            (course.teacher_name.includes(searchText) ||
                course.course_name.includes(searchText) ||
                course.credits.toString().includes(searchText) ||
                course.course_type.includes(searchText)) &&
            (courseType === '' || courseType === '全部' || course.course_type === courseType)
        );
        setFilteredData(filteredCourses);
        setCurrentPage(1); // 重置为第一页
    };

    const columns = [
        {
            title: '课程名',
            dataIndex: 'course_name',
        },
        {
            title: '学分',
            dataIndex: 'credits',
            sorter: (a, b) => a.credits - b.credits,
        },
        {
            title: '课程类型',
            dataIndex: 'course_type',
        },
        {
            title: '教师',
            dataIndex: 'teacher_name',
        },
        {
            title: '操作',
            key: 'action',
            render: () => (
                <Space size="middle">
                    <a onClick={() => handleDelete(selectedRowKeys)}>Delete</a>
                </Space>
            ),
        },
    ];

    const onSelectChange = (newSelectedRowKeys) => {
        setSelectedRowKeys(newSelectedRowKeys);
    };

    const rowSelection = {
        selectedRowKeys,
        onChange: onSelectChange,
    };

    const tableProps = {
        size: 'large',
        rowSelection,
        pagination: false,
    };

    const start = (currentPage - 1) * pageSize;
    const end = start + pageSize;
    const currentData = filteredData.slice(start, end);

    return (
        <div className="flex justify-center items-center min-h-screen bg-gray-100 p-64">
            <div className="w-3/4">
                <Card className="shadow-lg rounded-lg border border-gray-200">
                    <div className="mb-6 flex justify-between items-center">
                        <h2 className="text-xl font-semibold mx-auto">课程管理</h2>
                        <Button type="primary" icon={<PlusOutlined />} onClick={showAddCourseModal} className="mr-16">
                            添加课程
                        </Button>
                    </div>
                    <div className="flex mb-4">
                        <Input.Search
                            placeholder="输入教师名、课程名、学分或课程类型进行搜索"
                            allowClear
                            onSearch={handleSearch}
                            className="mr-32"
                        />
                        <Select
                            placeholder="选择课程类型"
                            allowClear
                            onChange={handleCourseTypeChange}
                            className="w-64 mr-16"
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
                    <Table
                        {...tableProps}
                        columns={columns}
                        dataSource={currentData}
                        rowKey="ID"
                    />
                    <div className="flex justify-center mt-4">
                        <Pagination
                            current={currentPage}
                            pageSize={pageSize}
                            total={filteredData.length}
                            onChange={setCurrentPage}
                            className="text-center"
                        />
                    </div>
                </Card>
            </div>

            <Modal title="添加课程" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                <Form form={form} layout="vertical" name="add_course_form">
                    <Form.Item
                        name="course_name"
                        label="课程名称"
                        rules={[{ required: true, message: '请输入课程名称' }]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name="credits"
                        label="学分"
                        rules={[{ required: true, message: '请输入学分' }]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name="course_type"
                        label="课程类型"
                        rules={[{ required: true, message: '请选择课程类型' }]}
                    >
                        <Select>
                            <Option value="专业课程">专业课程</Option>
                            <Option value="通识大类">通识大类</Option>
                            <Option value="创新思维与创业实践">创新思维与创业实践</Option>
                            <Option value="人文经典与文化传承">人文经典与文化传承</Option>
                            <Option value="艺术修养与审美体验">艺术修养与审美体验</Option>
                            <Option value="全球视野与文明对话">全球视野与文明对话</Option>
                            <Option value="科学探索与持续发展">科学探索与持续发展</Option>
                        </Select>
                    </Form.Item>
                    <Form.Item
                        name="teacher_name"
                        label="教师名称"
                        rules={[{ required: true, message: '请输入教师名称' }]}
                    >
                        <Input />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
};

export default App;